package interfaces

import "server/internal/pkg/models"

// EventDispatcher is an adapter interface to send one event through
// a single message channel
type EventDispatcher interface {
	Dispatch(event models.EventLike)
}
