package interfaces

import "server/internal/pkg/models"

type EventDispatcher interface {
	Dispatch(event models.EventLike)
}
