package cyclic_sender

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"time"
)

type CyclicSender struct {
	sender    interfaces.Sender
	lastEvent models.Event[models.RecordingEventData]
}

// interval time after which state updates are sent to clients
const interval = 100 * time.Millisecond

func NewCyclicSender(sender interfaces.Sender) *CyclicSender {
	return &CyclicSender{
		sender:    sender,
		lastEvent: models.NewRecordingIdleEvent(),
	}
}

func (n *CyclicSender) StartSendingPeriodicUpdates() {
	n.NotifyStopped()
	go func() {
		for {
			n.send()
			time.Sleep(interval)
		}
	}()
}

func (n *CyclicSender) NotifyStarted(res interfaces.StartedResponse) {
	n.lastEvent = models.NewRecordingRunningEvent(res.FileName, res.Started)
	n.send()
}

func (n *CyclicSender) NotifyStopped() {
	n.lastEvent = models.NewRecordingIdleEvent()
	n.send()
}

func (n *CyclicSender) send() {
	n.sender.Send(n.lastEvent)
}
