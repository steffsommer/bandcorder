package cyclic_sender

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"time"
)

type CyclicSender struct {
	dispatcher interfaces.EventDispatcher
	metaData   *interfaces.RecordingMetaData
}

// interval time after which state updates are sent to clients
const interval = 100 * time.Millisecond

func NewCyclicSender(dispatcher interfaces.EventDispatcher) *CyclicSender {
	return &CyclicSender{
		dispatcher: dispatcher,
	}
}

func (n *CyclicSender) StartSendingPeriodicUpdates() {
	n.NotifyStopped()
	go func() {
		for {
			n.dispatch()
			time.Sleep(interval)
		}
	}()
}

func (n *CyclicSender) NotifyStarted(res interfaces.RecordingMetaData) {
	n.metaData = &res
	n.dispatch()
}

func (n *CyclicSender) NotifyStopped() {
	n.metaData = nil
	n.dispatch()
}

func (n *CyclicSender) dispatch() {
	var event models.EventLike
	if n.metaData == nil {
		event = models.NewRecordingIdleEvent()
	} else {
		event = models.NewRecordingRunningEvent(n.metaData.FileName, n.metaData.Started)
	}
	n.dispatcher.Dispatch(event)
}
