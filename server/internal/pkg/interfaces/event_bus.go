package interfaces

import "server/internal/pkg/models"

type EventBus interface {
	EventDispatcher
	OnEvent(eventId models.EventId, cb func(any))
}
