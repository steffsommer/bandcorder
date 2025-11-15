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
		bpm:        initialBpm,
		dispatcher: dispatcher,
	}
}

func (m *MetronomeService) Start() {
	if m.ticker != nil {
		return
	}
	interval := time.Minute / time.Duration(m.bpm)
	m.ticker = time.NewTicker(interval)

	m.beat()
	go func() {
		for range m.ticker.C {
			m.beat()
		}
	}()
}

func (m *MetronomeService) beat() {
	event := models.NewMetronomeBeatEvent(m.beatCount)
	m.dispatcher.Dispatch(event)
	m.beatCount++
	fmt.Println(m.beatCount)
}

func (m *MetronomeService) Stop() {
	if m.ticker == nil {
		return
	}
	m.ticker.Stop()
	m.ticker = nil
	m.beatCount = 0
	event := models.NewMetronomeIdleEvent()
	go m.dispatcher.Dispatch(event)
}

func (m *MetronomeService) UpdateBpm(bpm int) {
	m.Stop()
	m.bpm = bpm
	m.beatCount = 0
	m.Start()
}
