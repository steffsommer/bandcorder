package interfaces

type Sender interface {
	Send(event string, data any)
}
