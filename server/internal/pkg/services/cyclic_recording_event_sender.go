package services

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"time"
)

type CyclicRecordingEventSender struct {
	eventBus         interfaces.EventBus
	targetDispatcher interfaces.EventDispatcher
	startedData      *models.RecordingStartedEventData
}

// interval time after which state updates are sent to clients
const interval = 100 * time.Millisecond

// Continously broadcast information about the current recording to all clients.
// Alleviates the need for time synchronization
func NewCyclicRecordingEventSender(
	sourceBus interfaces.EventBus,
	targetDispatcher interfaces.EventDispatcher,
) *CyclicRecordingEventSender {
	return &CyclicRecordingEventSender{
		eventBus:         sourceBus,
		targetDispatcher: targetDispatcher,
	}
}

func (n *CyclicRecordingEventSender) StartSendingPeriodicUpdates() {
	n.eventBus.OnEvent(models.RecordingStartedEvent, func(data any) {
		startedData, ok := data.(models.RecordingStartedEventData)
		if !ok {
			return
		}
		n.startedData = &startedData
		n.dispatch()
	})
	n.eventBus.OnEvent(models.RecordingStoppedEvent, func(_ any) {
		n.startedData = nil
		n.dispatch()
	})
	n.eventBus.OnEvent(models.RecordingAbortedEvent, func(_ any) {
		n.startedData = nil
		n.dispatch()
	})
	go func() {
		for {
			n.dispatch()
			time.Sleep(interval)
		}
	}()
}

func (n *CyclicRecordingEventSender) dispatch() {
	var event models.EventLike
	if n.startedData == nil {
		event = models.NewRecordingIdleEvent()
	} else {
		event = models.NewRecordingRunningEvent(n.startedData.FileName, n.startedData.Started)
	}
	n.targetDispatcher.Dispatch(event)
}
