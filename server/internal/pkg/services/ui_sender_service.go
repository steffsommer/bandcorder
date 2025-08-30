package services

import (
	"context"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type UiSenderService struct{
	appCtx context.Context
}

func NewUiSenderService(appCtx context.Context) *UiSenderService {
	return &UiSenderService{
		appCtx: appCtx,
	}
}

func (u *UiSenderService) Send(event string, data any) {
	log.Println("Sending event data")
	runtime.EventsEmit(u.appCtx, event, data)
}
