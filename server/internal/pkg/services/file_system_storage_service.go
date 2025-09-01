package services

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

type FileSystemStorageService struct {
	dir string
}

func NewFileSystemStorageService(dir string) *FileSystemStorageService {
	return &FileSystemStorageService{
		dir: dir,
	}
}

func (f *FileSystemStorageService) Save(fileName string, data []float32) error {
	path := filepath.Join(f.dir, fileName)
	file, err := os.Create(path)
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
