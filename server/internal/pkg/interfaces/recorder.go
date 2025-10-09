package interfaces

import "time"

type RecordingMetaData struct {
	FileName string
	Started  time.Time
}

// Recorder manages audio recording operations
type Recorder interface {
	// Start begins a new recording session and returns its metadata
	Start() (RecordingMetaData, error)

	// Stop ends the current recording session and saves the file
	Stop() error

	// Abort cancels the current recording session without saving
	Abort() error

	// GetMic returns the currently selected microphone device name
	GetMic() (string, error)
}
