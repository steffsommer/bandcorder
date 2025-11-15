package services

import (
	"fmt"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"time"
)

type MetronomeService struct {
	ticker     *time.Ticker
	bpm        int
	beatCount  int
	dispatcher interfaces.EventDispatcher
}

func NewMetronomeService(
	initialBpm int,
	dispatcher interfaces.EventDispatcher,
) *MetronomeService {
	return &MetronomeService{
		bpm: initialBpm,
	}
}

func (m *MetronomeService) Start() {
	if m.ticker != nil {
		return
	}
	interval := time.Minute / time.Duration(m.bpm)
	m.ticker = time.NewTicker(interval)
	for range m.ticker.C {
		fmt.Println("tick")
		event := models.NewMetronomeBeatEvent(m.beatCount)
		m.dispatcher.Dispatch(event)
		m.beatCount++
	}
}

func (m *MetronomeService) Stop() {
	if m.ticker == nil {
		return
	}
	m.ticker.Stop()
	m.ticker = nil
	event := models.NewMetronomeIdleEvent()
	m.dispatcher.Dispatch(event)
}

func (m *MetronomeService) UpdateBpm(bpm int) {
	m.bpm = bpm
	m.Start()
}
