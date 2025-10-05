package services

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
)

// BroadcastSender is an EventDispatcher, which writes to multiple other
// EventDispatchers
type BroadcastSender struct {
	senders []interfaces.EventDispatcher
}

func NewBroadcastSender(senders []interfaces.EventDispatcher) *BroadcastSender {
	return &BroadcastSender{
		senders: senders,
	}
}

func (b *BroadcastSender) Dispatch(event models.EventLike) {
	for _, sender := range b.senders {
		sender.Dispatch(event)
	}
}
