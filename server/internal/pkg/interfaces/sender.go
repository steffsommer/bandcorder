package interfaces

import "server/internal/pkg/models"

type Sender interface {
	Send(event models.EventLike)
}
