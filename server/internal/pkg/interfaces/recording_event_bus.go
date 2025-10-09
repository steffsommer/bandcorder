package interfaces

// RecordingEventBus provides an interface to notify all clients about
// recording state updates
type RecordingEventBus interface {
	NotifyStarted(res RecordingMetaData)
	NotifyStopped()
}
