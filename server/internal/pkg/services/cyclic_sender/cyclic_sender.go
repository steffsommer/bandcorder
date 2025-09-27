package cyclic_sender

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"time"
)

type CyclicSender struct {
	sender   interfaces.Sender
	metaData *interfaces.RecordingMetaData
}

// interval time after which state updates are sent to clients
const interval = 100 * time.Millisecond

func NewCyclicSender(sender interfaces.Sender) *CyclicSender {
	return &CyclicSender{
		sender: sender,
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

func (n *CyclicSender) NotifyStarted(res interfaces.RecordingMetaData) {
	n.metaData = &res
	n.send()
}

func (n *CyclicSender) NotifyStopped() {
	n.metaData = nil
	n.send()
}

func (n *CyclicSender) send() {
	var event models.EventLike
	if n.metaData == nil {
		event = models.NewRecordingIdleEvent()
	} else {
		event = models.NewRecordingRunningEvent(n.metaData.FileName, n.metaData.Started)
	}
	n.sender.Send(event)
}
