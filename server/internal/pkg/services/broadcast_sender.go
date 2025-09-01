package services

import "server/internal/pkg/interfaces"

type BroadcastSender struct {
	senders []interfaces.Sender
}

func NewBroadcastSender(senders []interfaces.Sender) *BroadcastSender {
	return &BroadcastSender{
		senders: senders,
	}
}

func (b *BroadcastSender) Send(event interfaces.EventID, data any) {
	for _, sender := range b.senders {
		sender.Send(event, data)
	}
}
