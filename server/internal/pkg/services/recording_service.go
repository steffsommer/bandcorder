package services

import (
	"errors"
	"fmt"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"server/internal/pkg/utils"
	"sync"
	"time"
	"unsafe"

	"github.com/gen2brain/malgo"
	"github.com/sirupsen/logrus"
)

// Recorder manages audio recording operations
type RecordingService struct {
	storageSerivce interfaces.StorageService
	audioProcessor interfaces.AudioProcessor
	eventBus       interfaces.EventBus
	ctx            *malgo.AllocatedContext
	device         *malgo.Device
	recording      []float32
	fileName       string
	done           chan bool
	mutex          sync.Mutex
}

// NewRecordingFacade creates a new RecorderService. It uses the default input device
func NewRecordingService(
	storageService interfaces.StorageService,
	audioProcessor interfaces.AudioProcessor,
	eventBus interfaces.EventBus,
) *RecordingService {
	return &RecordingService{
		done:           make(chan bool),
		recording:      make([]float32, 0, 1000),
		storageSerivce: storageService,
		audioProcessor: audioProcessor,
		eventBus:       eventBus,
	}
}

func (r *RecordingService) Init() error {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		return err
	}
	r.ctx = ctx
	return nil
}

func (r *RecordingService) Start() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isRunning() {
		return errors.New("Recording is already running")
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatF32
	deviceConfig.Capture.Channels = utils.Channels
	deviceConfig.SampleRate = utils.SampleRate
	deviceConfig.Alsa.NoMMap = 1

	onRecvFrames := func(pOutputSample, pInputSamples []byte, framecount uint32) {
		inputBuffer := make([]float32, framecount*utils.Channels)
		// Convert bytes to float32
		for i := range inputBuffer {
			idx := i * 4
			bits := uint32(pInputSamples[idx]) | uint32(pInputSamples[idx+1])<<8 |
				uint32(pInputSamples[idx+2])<<16 | uint32(pInputSamples[idx+3])<<24
			inputBuffer[i] = *(*float32)(unsafe.Pointer(&bits))
		}
		r.audioProcessor.Process(inputBuffer)
		r.recording = append(r.recording, inputBuffer...)
	}
	device, err := malgo.InitDevice(r.ctx.Context, deviceConfig, malgo.DeviceCallbacks{
		Data: onRecvFrames,
	})
	if err != nil {
		return fmt.Errorf("Failed to initialize device: %w", err)
	}
	if err := device.Start(); err != nil {
		return fmt.Errorf("Failed to start device: %w", err)
	}
	r.device = device
	r.fileName = fmt.Sprintf("recording-%s.wav", time.Now().Format("15-04-05"))
	ev := models.NewRecordingStartedEvent(r.fileName, time.Now())
	r.eventBus.Dispatch(ev)
	return nil
}

func (r *RecordingService) Stop() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if !r.isRunning() {
		return errors.New("No recording is running")
	}
	defer r.reset()
	ev := models.NewRecordingStoppedEvent()
	defer r.eventBus.Dispatch(ev)
	r.device.Uninit()
	if err := r.storageSerivce.Save(r.fileName, r.recording); err != nil {
		return err
	}
	logrus.Infof("Recorded %d samples (%.2f seconds)\n",
		len(r.recording), float64(len(r.recording))/float64(utils.SampleRate))
	return nil
}

func (r *RecordingService) Abort() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.isRunning() {
		return errors.New("No recording is running")
	}

	r.device.Uninit()
	r.reset()
	ev := models.NewRecordingAbortedEvent()
	r.eventBus.Dispatch(ev)
	return nil
}

func (r *RecordingService) reset() {
	r.fileName = ""
	r.recording = nil
}

func (r *RecordingService) isRunning() bool {
	return r.fileName != ""
}

func (r *RecordingService) GetMic() (string, error) {
	infos, err := r.ctx.Devices(malgo.Capture)
	if err != nil {
		return "", err
	}
	if len(infos) == 0 {
		return "", errors.New("No input devices found")
	}
	return infos[0].Name(), nil
}
