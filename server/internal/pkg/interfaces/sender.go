package interfaces

type EventID string

const (
	RunningEvent EventID = "RUNNING"
	IdleEvent            = "IDLE"
)

type Sender interface {
	Send(event EventID, data any)
}
