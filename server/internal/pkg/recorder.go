package recorder

import (
	"encoding/binary"
	"fmt"
	"os"
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
type Recorder struct {
	stream      *portaudio.Stream
	inputBuffer []float32
	recording   []float32
	isRunning   bool
	done        chan bool
}

var recorder *Recorder

func NewRecorder() (*Recorder, error) {
	if recorder != nil {
		return recorder, nil
	}
	recorder := Recorder{
		done:      make(chan bool),
		recording: make([]float32, 0, 1000),
	}
	return &recorder, nil
}

// Start starts a new recording
func (r *Recorder) Start() error {
	logrus.Info("Starting recording")
	if err := portaudio.Initialize(); err != nil {
		logrus.Fatalf("Failed to initialize PortAudio: %v", err)
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

	if err := r.stream.Start(); err != nil {
		logrus.Fatalf("Failed to start stream: %v", err)
	}

	go func() {
		for {
			select {
			case <-r.done:
				return
			default:
				if err := r.stream.Read(); err != nil {
					logrus.Printf("Error reading from stream: %v", err)
					continue
				}
				r.recording = append(r.recording, r.inputBuffer...)
			}
		}
	}()

	logrus.Info("Recording started")
	return nil
}

// Stop finishes the current recording
func (r *Recorder) Stop() error {
	logrus.Info("Stopping recording")
	if err := r.stream.Stop(); err != nil {
		logrus.Printf("Error stopping stream: %v", err)
	}
	r.done <- true

	filename := fmt.Sprintf("recording_%d.raw", time.Now().Unix())
	if err := saveAudioFile(r.recording, "test_file.wav"); err != nil {
		logrus.Fatalf("Failed to save audio to file: %v", err)
	}

	logrus.Infof("Created recording %s, duration: %.2f seconds\n",
		filename, float64(len(r.recording))/float64(sampleRate))

	defer r.stream.Close()
	logrus.Info("Recording stopped")
	return nil
}

// Abort aborts the current recording without saving it
func (r *Recorder) Abort() error {
	r.done <- true
	r.stream.Stop()
	r.stream.Close()
	r.isRunning = false
	return nil
}

// for a description for the WAV format, see http://soundfile.sapp.org/doc/WaveFormat/
func saveAudioFile(data []float32, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// WAV header
	dataSize := len(data) * 4 // 4 bytes per float32
	fileSize := 36 + dataSize // 36 = header length - 4 (RIFF) - 4 (filesize)

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

	for _, sample := range data {
		if err := binary.Write(file, binary.LittleEndian, sample); err != nil {
			return fmt.Errorf("failed to write sample: %w", err)
		}
	}

	return nil
}
