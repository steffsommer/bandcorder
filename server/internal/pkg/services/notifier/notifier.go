package notifier

import (
	"server/internal/pkg/interfaces"
	"time"
)

type Notifier struct {
	sender    interfaces.Sender
	lastEvent struct {
		id   interfaces.EventID
		data any
	}
}

// interval time after which state updates are sent to clients
const interval = 100 * time.Millisecond

func NewNotifier(sender interfaces.Sender) *Notifier {
	return &Notifier{
		sender: sender,
	}
}

func (n *Notifier) StartSendingPeriodicUpdates() {
	go func() {
		n.lastEvent.id = interfaces.RunningEvent
		n.lastEvent.data = recordingRunningEvent{
			FileName: "test-file-name.wav",
			Started:  time.Now(),
		}
		for {
			n.send()
			time.Sleep(interval)
		}
	}()
}

func (n *Notifier) NotifyStarted() {
	n.lastEvent.id = interfaces.RunningEvent
	n.lastEvent.data = recordingRunningEvent{
		FileName: "TODO",
		Started:  time.Now(),
	}
	n.send()
}

func (n *Notifier) NotifyStopped() {
	n.lastEvent.id = interfaces.IdleEvent
}

func (n *Notifier) send() {
	if n.lastEvent.id == "" {
		return
	}
	n.sender.Send(n.lastEvent.id, n.lastEvent.data)
}
