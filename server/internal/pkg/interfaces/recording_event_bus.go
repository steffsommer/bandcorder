package interfaces

// ClientNotifier sends state changes to the client
// It should send a state change as soon as a state change happens
// and every 100ms.
// The data should be sent as JSON
type RecordingEventBus interface {
	NotifyStarted()
	NotifyStopped()
}
