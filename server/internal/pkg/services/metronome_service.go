package services

import (
	"errors"
	"fmt"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"sync"
	"time"
)

const (
	minBpm = 40
	maxBpm = 240
)

type MetronomeService struct {
	ticker          *time.Ticker
	bpm             int
	beatCount       int
	dispatcher      interfaces.EventDispatcher
	playbackService interfaces.PlaybackService
	mutex           sync.Mutex
}

func NewMetronomeService(
	initialBpm int,
	dispatcher interfaces.EventDispatcher,
	playbackService interfaces.PlaybackService,
) *MetronomeService {
	return &MetronomeService{
		bpm:             initialBpm,
		dispatcher:      dispatcher,
		playbackService: playbackService,
	}
}

func (m *MetronomeService) Start() error {
	// m.mutex.Lock()
	// defer m.mutex.Unlock()
	if m.ticker != nil {
		return errors.New("Metronome is already running")
	}
	interval := time.Minute / time.Duration(m.bpm)
	m.ticker = time.NewTicker(interval)

	m.beat()
	go func() {
		for range m.ticker.C {
			m.beat()
		}
	}()
	return nil
}

func (m *MetronomeService) beat() {
	event := models.NewMetronomeBeatEvent(m.beatCount)
	m.dispatcher.Dispatch(event)
	m.beatCount++
	m.playbackService.Play(interfaces.MetronomeClick)
}

func (m *MetronomeService) Stop() error {
	// m.mutex.Lock()
	// defer m.mutex.Unlock()
	if m.ticker == nil {
		return errors.New("Metronome is not running")
	}
	m.ticker.Stop()
	m.ticker = nil
	m.beatCount = 0
	event := models.NewMetronomeIdleEvent()
	go m.dispatcher.Dispatch(event)
	return nil
}

func (m *MetronomeService) UpdateBpm(bpm int) error {
	// m.mutex.Lock()
	// defer m.mutex.Unlock()
	if bpm < minBpm || bpm > maxBpm {
		return fmt.Errorf("BPM must be in the range %d-%d", minBpm, maxBpm)
	}
	m.Stop()
	m.bpm = bpm
	m.beatCount = 0
	m.Start()
	return nil
}
