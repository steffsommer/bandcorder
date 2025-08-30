package services

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/labstack/gommon/log"
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
	stream      *portaudio.Stream
	inputBuffer []float32
	recording   []float32
	isRunning   bool
	done        chan bool
	mutex       sync.Mutex
}

func NewRecorderService() *RecorderService {
	return &RecorderService{
		done:      make(chan bool),
		recording: make([]float32, 0, 1000),
	}
}

// Init should be called once to initialize the underlying audio system
// for example, init scans for available audio devices
func (r *RecorderService) Init() error {
	return portaudio.Initialize()
}

// Start starts a new recording. The recording will fill an in-memory buffer
// until either Stop() or Abort() are called
func (r *RecorderService) Start() error {
	logrus.Info("Starting recording")
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.isRunning {
		return errors.New("Recording is already running")
	}

	inputDevice, err := portaudio.DefaultInputDevice()
	if err != nil {
		logrus.Fatalf("Failed to get default input device: %v", err)
	}
	logrus.Infof("Recording from: %s\n", inputDevice.Name)

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
		logrus.Fatalf("Failed to open stream: %v", err)
	}
	r.stream = stream

	log.Info("Starting stream")
	if err := r.stream.Start(); err != nil {
		logrus.Fatalf("Failed to start stream: %v", err)
	}

	go func() {
		for {
			select {
			case <-r.done:
				return
			default:
				// log.Info("reading from stream")
				if err := r.stream.Read(); err != nil {
					logrus.Printf("Error reading from stream: %v", err)
					continue
				}
				// log.Info("done reading from stream")
				r.recording = append(r.recording, r.inputBuffer...)
			}
		}
	}()

	logrus.Info("Recording started")
	r.isRunning = true
	return nil
}

// Stop stops the current recording and writes the recorded audio to a
// wav file
func (r *RecorderService) Stop() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.done <- true

	logrus.Info("Closing the stream")
	if err := r.stream.Close(); err != nil {
		logrus.Printf("Error closing stream: %v", err)
	}

	filename := fmt.Sprintf("recording_%d.raw", time.Now().Unix())
	if err := saveWavFile(r.recording, "test_file.wav"); err != nil {
		logrus.Fatalf("Failed to save audio to file: %v", err)
	}

	fmt.Printf("Recording saved to %s\n", filename)
	fmt.Printf("Recorded %d samples (%.2f seconds)\n",
		len(r.recording), float64(len(r.recording))/float64(sampleRate))

	logrus.Info("Recording stopped")
	r.isRunning = false
	r.recording = nil
	return nil
}

// Abort aborts the current recording without saving it
func (r *RecorderService) Abort() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.stream.Stop()
	r.done <- true
	r.stream.Close()
	r.isRunning = false
	return nil
}

func saveWavFile(data []float32, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// WAV header
	dataSize := len(data) * 4 // 4 bytes per float32
	fileSize := 36 + dataSize

	// Write WAV header
	file.WriteString("RIFF")
	binary.Write(file, binary.LittleEndian, uint32(fileSize))
	file.WriteString("WAVE")
	file.WriteString("fmt ")
	binary.Write(file, binary.LittleEndian, uint32(16))                    // PCM format size
	binary.Write(file, binary.LittleEndian, uint16(3))                     // IEEE float format
	binary.Write(file, binary.LittleEndian, uint16(channels))              // Number of channels
	binary.Write(file, binary.LittleEndian, uint32(sampleRate))            // Sample rate
	binary.Write(file, binary.LittleEndian, uint32(sampleRate*channels*4)) // Byte rate
	binary.Write(file, binary.LittleEndian, uint16(channels*4))            // Block align
	binary.Write(file, binary.LittleEndian, uint16(32))                    // Bits per sample
	file.WriteString("data")
	binary.Write(file, binary.LittleEndian, uint32(dataSize))

	// Write audio data
	for _, sample := range data {
		if err := binary.Write(file, binary.LittleEndian, sample); err != nil {
			return fmt.Errorf("failed to write sample: %w", err)
		}
	}

	return nil
}
