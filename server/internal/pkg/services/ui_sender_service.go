package services

import (
	"context"
	"server/internal/pkg/models"

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

func (u *UiSenderService) Send(event models.EventLike) {
	if u.wailsCtx == nil {
		panic("Wails Context has not been set")
	}
	runtime.EventsEmit(u.wailsCtx, string(event.GetId()), event.GetData())
}
