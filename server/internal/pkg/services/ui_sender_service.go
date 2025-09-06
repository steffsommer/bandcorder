package services

import (
	"context"
	"server/internal/pkg/interfaces"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type UiSenderService struct {
	wailsCtx context.Context
}

func NewUiSenderService() *UiSenderService {
	return &UiSenderService{}
}

func (u *UiSenderService) Init(appCtx context.Context) {
	u.wailsCtx = appCtx
}

func (u *UiSenderService) Send(event interfaces.EventID, data any) {
	if u.wailsCtx == nil {
		panic("Wails Context has not been set")
	}
	runtime.EventsEmit(u.wailsCtx, string(event), data)
}
