package services

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

const bytesPerSample = 4

type FileSystemStorageService struct {
	dir          string
	channelCount int
	sampleRate   int
}

func NewFileSystemStorageService(dir string, channelCount int, sampleRate int) *FileSystemStorageService {
	return &FileSystemStorageService{
		dir:          dir,
		channelCount: channelCount,
		sampleRate:   sampleRate,
	}
}

func (f *FileSystemStorageService) Save(fileName string, data []float32) error {
	path := filepath.Join(f.dir, fileName)
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Use buffered writer for much better performance
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// WAV header
	dataSize := len(data) * 4 // 4 bytes per float32
	fileSize := 36 + dataSize

	// Write WAV header
	writer.WriteString("RIFF")
	binary.Write(writer, binary.LittleEndian, uint32(fileSize))
	writer.WriteString("WAVE")
	writer.WriteString("fmt ")
	binary.Write(writer, binary.LittleEndian, uint32(16))           // PCM format size
	binary.Write(writer, binary.LittleEndian, uint16(3))            // IEEE float format
	binary.Write(writer, binary.LittleEndian, uint16(1))            // Number of channels (1 for mono)
	binary.Write(writer, binary.LittleEndian, uint32(f.sampleRate)) // Sample rate in Hz, e.g. 44100 for 44.1kHz
	binary.Write(writer, binary.LittleEndian,
		uint32(f.sampleRate*f.channelCount*bytesPerSample)) // Byte rate
	binary.Write(writer, binary.LittleEndian, uint16(1*4)) // Block align
	binary.Write(writer, binary.LittleEndian, uint16(32))  // Bits per sample

	writer.WriteString("data")
	binary.Write(writer, binary.LittleEndian, uint32(dataSize))

	// Write ALL audio data in one operation - this is the key fix!
	return binary.Write(writer, binary.LittleEndian, data)
}
