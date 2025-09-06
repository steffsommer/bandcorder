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
	n.NotifyStopped()
	go func() {
		for {
			n.send()
			time.Sleep(interval)
		}
	}()
}

func (n *EventBus) NotifyStarted(res interfaces.StartedResponse) {
	n.lastEvent.id = interfaces.RecordingStateEvent
	n.lastEvent.data = recordingRunningEvent{
		State:    RUNNING,
		FileName: res.FileName,
		Started:  res.Started,
	}
	n.send()
}

func (n *EventBus) NotifyStopped() {
	n.lastEvent.id = interfaces.RecordingStateEvent
	n.lastEvent.data = recordingRunningEvent{
		State: IDLE,
	}
	n.send()
}

func (n *EventBus) send() {
	n.sender.Send(n.lastEvent.id, n.lastEvent.data)
}
