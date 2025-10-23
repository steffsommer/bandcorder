package services

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const bytesPerSample = 4

// WAVHeader represents the structure of a WAV file header
type WAVHeader struct {
	ChunkID       [4]byte // "RIFF"
	ChunkSize     uint32  // File size - 8
	Format        [4]byte // "WAVE"
	Subchunk1ID   [4]byte // "fmt "
	Subchunk1Size uint32  // Size of format chunk
	AudioFormat   uint16  // Audio format (1 = PCM)
	NumChannels   uint16  // Number of channels
	SampleRate    uint32  // Sample rate
	ByteRate      uint32  // Byte rate
	BlockAlign    uint16  // Block align
	BitsPerSample uint16  // Bits per sample
	Subchunk2ID   [4]byte // "data"
	Subchunk2Size uint32  // Size of data chunk
}

// FileSystemStorageService creates WAV Files from raw float32 audio data and saves
// them to the file system. For each day with recordings, a folder with the name of
// the current date is created.
type FileSystemStorageService struct {
	baseDir      string
	channelCount int
	sampleRate   int
	timeProvider interfaces.TimeProvider
}

func NewFileSystemStorageService(
	dir string,
	channelCount int,
	sampleRate int,
	timeProvider interfaces.TimeProvider,
) *FileSystemStorageService {
	return &FileSystemStorageService{
		baseDir:      dir,
		channelCount: channelCount,
		sampleRate:   sampleRate,
		timeProvider: timeProvider,
	}
}

func (f *FileSystemStorageService) Save(fileName string, data []float32) error {
	targetDir, err := f.getTargetDirCreateIfNeeded()
	if err != nil {
		return err
	}
	path := filepath.Join(targetDir, fileName)
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
	err = binary.Write(writer, binary.LittleEndian, data)
	logrus.Infof("Successfully wrote audio file %s", path)
	return err
}

func (f *FileSystemStorageService) getTargetDirCreateIfNeeded() (string, error) {
	today := f.timeProvider.Now()
	absTargetDir := f.getAbsDateDir(today)
	var err error
	if _, err = os.Stat(absTargetDir); os.IsNotExist(err) {
		err = os.Mkdir(absTargetDir, 0755)
	}
	if err != nil {
		return "", err
	}
	return absTargetDir, err
}

func (f *FileSystemStorageService) getAbsDateDir(date time.Time) string {
	dateDir := date.Format("2006-01-02")
	return filepath.Join(f.baseDir, dateDir)
}

func (f *FileSystemStorageService) GetRecordings(
	date time.Time,
) ([]models.RecordingInfo, error) {
	infos := make([]models.RecordingInfo, 0)
	absDir := f.getAbsDateDir(date)
	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		return infos, nil
	}
	entries, err := os.ReadDir(absDir)
	if os.IsNotExist(err) {
		return infos, nil
	} else if err != nil {
		return infos, err
	}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".wav") {
			continue
		}
		recAbsPath := filepath.Join(absDir, entry.Name())
		duration, err := getWavDuration(recAbsPath)
		if err != nil {
			return infos, err
		}
		info := models.RecordingInfo{
			FileName:        entry.Name(),
			DurationSeconds: duration,
		}
		infos = append(infos, info)
	}
	slices.SortFunc(infos, func(a, b models.RecordingInfo) int {
		if a.FileName < b.FileName {
			return 1
		} else {
			return -1
		}
	})
	return infos, nil
}

func getWavDuration(absPath string) (uint32, error) {
	file, err := os.Open(absPath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var header WAVHeader
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return 0, fmt.Errorf("failed to read WAV header: %w", err)
	}

	// Verify it's a valid WAV file
	if string(header.ChunkID[:]) != "RIFF" || string(header.Format[:]) != "WAVE" {
		return 0, fmt.Errorf("not a valid WAV file")
	}

	// Handle cases where fmt chunk might not be exactly 16 bytes
	if header.Subchunk1Size != 16 {
		// Seek to the data chunk
		_, err = file.Seek(20+int64(header.Subchunk1Size), 0)
		if err != nil {
			return 0, fmt.Errorf("failed to seek to data chunk: %w", err)
		}

		// Read data chunk header
		err = binary.Read(file, binary.LittleEndian, &header.Subchunk2ID)
		if err != nil {
			return 0, fmt.Errorf("failed to read data chunk ID: %w", err)
		}

		err = binary.Read(file, binary.LittleEndian, &header.Subchunk2Size)
		if err != nil {
			return 0, fmt.Errorf("failed to read data chunk size: %w", err)
		}
	}

	// Verify we found the data chunk
	if string(header.Subchunk2ID[:]) != "data" {
		return 0, fmt.Errorf("data chunk not found")
	}

	// Calculate duration: data size / byte rate
	if header.ByteRate == 0 {
		return 0, fmt.Errorf("invalid byte rate (zero)")
	}

	durationSeconds := float64(header.Subchunk2Size) / float64(header.ByteRate)

	return uint32(durationSeconds), nil
}

func (f *FileSystemStorageService) RenameRecording(
	oldFileName string,
	newFileName string,
	date time.Time,
) error {
	absDateDir := f.getAbsDateDir(date)
	if _, err := os.Stat(absDateDir); os.IsNotExist(err) {
		return fmt.Errorf("No recordings exist for given date %s", date)
	}
	absFileOld := filepath.Join(absDateDir, oldFileName)
	if _, err := os.Stat(absFileOld); os.IsNotExist(err) {
		return fmt.Errorf("Recording file %s does not exist for date %s", oldFileName, date)
	}
	absFileNew := filepath.Join(absDateDir, newFileName)
	return os.Rename(absFileOld, absFileNew)
}

func (f *FileSystemStorageService) DeleteRecording(fileName string, date time.Time) error {
	absDateDir := f.getAbsDateDir(date)
	if _, err := os.Stat(absDateDir); os.IsNotExist(err) {
		return fmt.Errorf("No recordings exist for given date %s", date)
	}
	absFile := filepath.Join(absDateDir, fileName)
	if _, err := os.Stat(absFile); os.IsNotExist(err) {
		return fmt.Errorf("Recording file %s does not exist for date %s", fileName, date)
	}
	return os.Remove(absFile)
}

func (f *FileSystemStorageService) RenameLastRecording(fileName string) error {
	entries, err := os.ReadDir(f.baseDir)
	if err != nil {
		return err
	}

	var dirs []os.DirEntry
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry)
		}
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name() > dirs[j].Name()
	})

	for _, dir := range dirs {
		dirPath := filepath.Join(f.baseDir, dir.Name())
		files, err := os.ReadDir(dirPath)
		if err != nil {
			continue
		}

		var fileNames []string
		for _, file := range files {
			if !file.IsDir() {
				fileNames = append(fileNames, file.Name())
			}
		}

		if len(fileNames) == 0 {
			continue
		}

		sort.Strings(fileNames)

		oldPath := filepath.Join(dirPath, fileNames[0])
		newPath := filepath.Join(dirPath, fileName)
		return os.Rename(oldPath, newPath)
	}

	return fmt.Errorf("No recording exist")
}
