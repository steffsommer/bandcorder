package services

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"time"
)

type CyclicRecordingEventSender struct {
	dispatcher interfaces.EventDispatcher
	metaData   *interfaces.RecordingMetaData
}

// interval time after which state updates are sent to clients
const interval = 100 * time.Millisecond

func NewCyclicRecordingEventSender(dispatcher interfaces.EventDispatcher) *CyclicRecordingEventSender {
	return &CyclicRecordingEventSender{
		dispatcher: dispatcher,
	}
}

func (n *CyclicRecordingEventSender) StartSendingPeriodicUpdates() {
	n.NotifyStopped()
	go func() {
		for {
			n.dispatch()
			time.Sleep(interval)
		}
	}()
}

func (n *CyclicRecordingEventSender) NotifyStarted(res interfaces.RecordingMetaData) {
	n.metaData = &res
	n.dispatch()
}

func (n *CyclicRecordingEventSender) NotifyStopped() {
	n.metaData = nil
	n.dispatch()
}

func (n *CyclicRecordingEventSender) dispatch() {
	var event models.EventLike
	if n.metaData == nil {
		event = models.NewRecordingIdleEvent()
	} else {
		event = models.NewRecordingRunningEvent(n.metaData.FileName, n.metaData.Started)
	}
	n.dispatcher.Dispatch(event)
}
