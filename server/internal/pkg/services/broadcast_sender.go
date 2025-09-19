package services

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
)

type BroadcastSender struct {
	senders []interfaces.Sender
}

func NewBroadcastSender(senders []interfaces.Sender) *BroadcastSender {
	return &BroadcastSender{
		senders: senders,
	}
}

func (b *BroadcastSender) Send(event models.EventLike) {
	for _, sender := range b.senders {
		sender.Send(event)
	}
}
