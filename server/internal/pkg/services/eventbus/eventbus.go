package eventbus

import (
	"server/internal/pkg/interfaces"
	"time"
)

type EventBus struct {
	sender    interfaces.Sender
	lastEvent struct {
		id   interfaces.EventID
		data any
	}
}

// interval time after which state updates are sent to clients
const interval = 100 * time.Millisecond

func NewEventBus(sender interfaces.Sender) *EventBus {
	return &EventBus{
		sender: sender,
	}
}

func (n *EventBus) StartSendingPeriodicUpdates() {
	go func() {
		// n.lastEvent.id = interfaces.RunningEvent
		// n.lastEvent.data = recordingRunningEvent{
		// 	FileName: "test-file-name.wav",
		// 	Started:  time.Now(),
		// }
		for {
			n.send()
			time.Sleep(interval)
		}
	}()
}

func (n *EventBus) NotifyStarted() {
	n.lastEvent.id = interfaces.RecordingStateEvent
	n.lastEvent.data = recordingRunningEvent{
		State:    RUNNING,
		FileName: "TODO",
		Started:  time.Now(),
	}
	n.send()
}

func (n *EventBus) NotifyStopped() {
	n.lastEvent.id = interfaces.RecordingStateEvent
	n.lastEvent.data = recordingRunningEvent{
		State: IDLE,
	}
}

func (n *EventBus) send() {
	if n.lastEvent.id == "" {
		return
	}
	n.sender.Send(n.lastEvent.id, n.lastEvent.data)
}
