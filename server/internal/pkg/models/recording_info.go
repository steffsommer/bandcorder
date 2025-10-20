package models

type RecordingInfo struct {
	FileName        string    `json:"fileName"`
	DurationSeconds uint32    `json:"durationSeconds"`
}
