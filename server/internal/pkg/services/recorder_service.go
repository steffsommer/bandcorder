package services

import (
	"errors"
	"fmt"
	"server/internal/pkg/interfaces"
	"sync"
	"time"
	"unsafe"

	"github.com/gen2brain/malgo"
	"github.com/sirupsen/logrus"
)

const (
	sampleRate = 44100
	channels   = 1
	bufferSize = 1024
)

type RecorderService struct {
	storageSerivce interfaces.StorageService
	audioProcessor interfaces.AudioProcessor
	ctx            *malgo.AllocatedContext
	device         *malgo.Device
	recording      []float32
	fileName       string
	done           chan bool
	mutex          sync.Mutex
}

func NewRecorderService(
	storageService interfaces.StorageService,
	audioProcessor interfaces.AudioProcessor,
) *RecorderService {
	return &RecorderService{
		done:           make(chan bool),
		recording:      make([]float32, 0, 1000),
		storageSerivce: storageService,
		audioProcessor: audioProcessor,
	}
}

func (r *RecorderService) Init() error {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		return err
	}
	r.ctx = ctx
	return nil
}

func (r *RecorderService) Start() (interfaces.RecordingMetaData, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isRunning() {
		return interfaces.RecordingMetaData{}, errors.New("Recording is already running")
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatF32
	deviceConfig.Capture.Channels = channels
	deviceConfig.SampleRate = sampleRate
	deviceConfig.Alsa.NoMMap = 1

	onRecvFrames := func(pOutputSample, pInputSamples []byte, framecount uint32) {
		inputBuffer := make([]float32, framecount*channels)
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
		return interfaces.RecordingMetaData{}, fmt.Errorf("Failed to initialize device: %w", err)
	}

	if err := device.Start(); err != nil {
		return interfaces.RecordingMetaData{}, fmt.Errorf("Failed to start device: %w", err)
	}

	r.device = device
	r.fileName = fmt.Sprintf("recording-%s.wav", time.Now().Format("15-04-05"))

	return interfaces.RecordingMetaData{
		FileName: r.fileName,
		Started:  time.Now(),
	}, nil
}

func (r *RecorderService) Stop() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.isRunning() {
		return errors.New("No recording is running")
	}

	r.device.Uninit()

	if err := r.storageSerivce.Save(r.fileName, r.recording); err != nil {
		return err
	}

	logrus.Infof("Recorded %d samples (%.2f seconds)\n",
		len(r.recording), float64(len(r.recording))/float64(sampleRate))

	r.reset()
	return nil
}

func (r *RecorderService) Abort() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.isRunning() {
		return errors.New("No recording is running")
	}

	r.device.Uninit()
	r.reset()
	return nil
}

func (r *RecorderService) reset() {
	r.fileName = ""
	r.recording = nil
}

func (r *RecorderService) isRunning() bool {
	return r.fileName != ""
}

func (r *RecorderService) GetMic() (string, error) {
	infos, err := r.ctx.Devices(malgo.Capture)
	if err != nil {
		return "", err
	}
	if len(infos) == 0 {
		return "", errors.New("No input devices found")
	}
	return infos[0].Name(), nil
}
