package interfaces

type EventID string

const (
	RecordingStateEvent EventID = "RecordingState"
)

type Sender interface {
	Send(event EventID, data any)
}
