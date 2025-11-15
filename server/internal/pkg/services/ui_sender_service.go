package services

import (
	"context"
	"server/internal/pkg/models"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// UiSenderService is an EventDispatcher, which sends Events to the Desktop Ui
// via the Event System built into wails. This is also a websocket connection
// with different requirements related to timeouts etc. than the outgoing
// Websocket connections with the app clients
type UiSenderService struct {
	wailsCtx context.Context
}

func NewUiSenderService() *UiSenderService {
	return &UiSenderService{}
}

func (u *UiSenderService) Init(appCtx context.Context) {
	u.wailsCtx = appCtx
}

func (u *UiSenderService) Dispatch(event models.EventLike) {
	if u.wailsCtx == nil {
		panic("Wails Context has not been set")
	}
	// dispatching events in a separate Goroutine is crucial, otherwise GTK may crash
	go runtime.EventsEmit(u.wailsCtx, string(event.GetId()), event.GetData())
}
