package services

import (
	"errors"
	"fmt"
	"server/internal/pkg/interfaces"
	"sync"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/sirupsen/logrus"
)

const (
	sampleRate = 44100
	channels   = 1 // mono
	bufferSize = 1024
)

// Recorder creates audio recordings
//
// TODOs:
// - handle microphone selection
type RecorderService struct {
	storageSerivce interfaces.StorageService
	stream         *portaudio.Stream
	inputBuffer    []float32
	recording      []float32
	fileName       string
	done           chan bool
	mutex          sync.Mutex
}

func NewRecorderService(storageService interfaces.StorageService) *RecorderService {
	return &RecorderService{
		done:           make(chan bool),
		recording:      make([]float32, 0, 1000),
		storageSerivce: storageService,
	}
}

// Init should be called once to initialize the underlying audio system
// for example, init scans for available audio devices
func (r *RecorderService) Init() error {
	return portaudio.Initialize()
}

// Start starts a new recording. The recording will fill an in-memory buffer
// until either Stop() or Abort() are called
func (r *RecorderService) Start() (interfaces.StartedResponse, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isRunning() {
		return interfaces.StartedResponse{}, errors.New("Recording is already running")
	}

	inputDevice, err := portaudio.DefaultInputDevice()
	if err != nil {
		fmt.Errorf("Failed to get default input device: %w", err)
	}

	inputParams := portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{
			Device:   inputDevice,
			Channels: channels,
			Latency:  inputDevice.DefaultLowInputLatency,
		},
		SampleRate:      sampleRate,
		FramesPerBuffer: bufferSize,
	}
	r.inputBuffer = make([]float32, bufferSize)

	stream, err := portaudio.OpenStream(inputParams, r.inputBuffer)
	if err != nil {
		return interfaces.StartedResponse{}, fmt.Errorf("Failed to open stream: %w", err)
	}
	r.stream = stream

	if err := r.stream.Start(); err != nil {
		return interfaces.StartedResponse{}, fmt.Errorf("Failed to start stream: %w", err)
	}

	go func() {
		for {
			select {
			case <-r.done:
				return
			default:
				if err := r.stream.Read(); err != nil {
					logrus.Errorf("Error reading from stream: %v", err)
					continue
				}
				r.recording = append(r.recording, r.inputBuffer...)
			}
		}
	}()

	r.fileName = fmt.Sprintf("recording_%d.wav", time.Now().Unix())
	return interfaces.StartedResponse{
		FileName: r.fileName,
		Started:  time.Now(),
	}, nil
}

// Stop stops the current recording and writes the recorded audio to a wav file
func (r *RecorderService) Stop() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.isRunning() {
		return errors.New("No recording is running")
	}

	r.done <- true

	if err := r.stream.Close(); err != nil {
		return err
	}

	if err := r.storageSerivce.Save(r.fileName, r.recording); err != nil {
		return err
	}

	logrus.Infof("Recorded %d samples (%.2f seconds)\n",
		len(r.recording), float64(len(r.recording))/float64(sampleRate))

	r.reset()
	return nil
}

// Abort aborts the current recording without saving it
func (r *RecorderService) Abort() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.isRunning() {
		return errors.New("No recording is running")
	}

	r.done <- true
	err := r.stream.Close()
	if err != nil {
		return err
	}
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
