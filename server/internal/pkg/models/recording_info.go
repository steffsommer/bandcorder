package models

import "time"

type RecordingInfo struct {
	FileName        string    `json:"fileName"`
	DurationSeconds uint32    `json:"durationSeconds"`
	ModTime         time.Time `json:"modTime"`
}
