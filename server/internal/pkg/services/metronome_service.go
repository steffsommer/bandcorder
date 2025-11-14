package services

import (
	"fmt"
	"time"
)

type MetronomeService struct {
	ticker *time.Ticker
	bpm    int
}

func NewMetronomeService(initialBpm int) *MetronomeService {
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
	}
}

func (m *MetronomeService) Stop() {
	if m.ticker == nil {
		return
	}
	m.ticker.Stop()
	m.ticker = nil
}

func (m *MetronomeService) UpdateBpm(bpm int) {
	m.bpm = bpm
	m.Start()
}
