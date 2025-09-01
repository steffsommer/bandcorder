package services

import (
	"context"
	"server/internal/pkg/interfaces"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type UiSenderService struct {
	appCtx context.Context
}

func NewUiSenderService(appCtx context.Context) *UiSenderService {
	return &UiSenderService{
		appCtx: appCtx,
	}
}

func (u *UiSenderService) Send(event interfaces.EventID, data any) {
	runtime.EventsEmit(u.appCtx, string(event), data)
}
