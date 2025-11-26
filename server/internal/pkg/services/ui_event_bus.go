package services

import (
	"context"
	"server/internal/pkg/models"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// UiEventBus is capable of sending events to and receiving them from the
// frontend. It also may receive events emitted within the backend, making
// it a general purpose event bus
// The BE-BE communication is built on top of the wails-specific event system
type UiEventBus struct {
	wailsCtx    context.Context
	subscribers map[models.EventId][]func(any)
	mu          sync.RWMutex
}

func NewUiEventBus() *UiEventBus {
	return &UiEventBus{
		subscribers: make(map[models.EventId][]func(any)),
	}
}

func (u *UiEventBus) Init(appCtx context.Context) {
	u.wailsCtx = appCtx
}

func (u *UiEventBus) Dispatch(event models.EventLike) {
	if u.wailsCtx == nil {
		panic("Wails Context has not been set")
	}

	go runtime.EventsEmit(u.wailsCtx, string(event.GetId()), event.GetData())

	u.mu.RLock()
	callbacks := u.subscribers[event.GetId()]
	u.mu.RUnlock()

	for _, cb := range callbacks {
		go cb(event.GetData())
	}
}

func (u *UiEventBus) OnEvent(eventId models.EventId, cb func(any)) {
	u.mu.Lock()
	u.subscribers[eventId] = append(u.subscribers[eventId], cb)
	u.mu.Unlock()
}
