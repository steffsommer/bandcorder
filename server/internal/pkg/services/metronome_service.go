package services

import (
	"errors"
	"fmt"
	"math"
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
	m.mutex.Lock()
	if m.ticker != nil {
		m.mutex.Unlock()
		return errors.New("Metronome is already running")
	}
	m.startInternal()
	event := models.NewMetronomeStateChangeEvent(true, m.bpm)
	m.mutex.Unlock()
	go m.dispatcher.Dispatch(event)
	return nil
}

func (m *MetronomeService) beat() {
	beatCount := m.beatCount
	m.beatCount = (m.beatCount + 1) % (math.MaxInt - 1)
	event := models.NewMetronomeBeatEvent(beatCount)
	m.dispatcher.Dispatch(event)
	m.playbackService.Play(interfaces.MetronomeClick)
}
func (m *MetronomeService) Stop() error {
	m.mutex.Lock()
	if m.ticker == nil {
		return errors.New("Metronome is not running")
	}
	m.stopInternal()
	m.mutex.Unlock()
	event := models.NewMetronomeStateChangeEvent(false, m.bpm)
	go m.dispatcher.Dispatch(event)
	return nil
}

func (m *MetronomeService) UpdateBpm(bpm int) error {
	m.mutex.Lock()
	if bpm < minBpm || bpm > maxBpm {
		m.mutex.Unlock()
		return fmt.Errorf("BPM must be in the range %d-%d", minBpm, maxBpm)
	}

	m.bpm = bpm
	isRunning := m.ticker != nil
	if isRunning {
		m.stopInternal()
		m.startInternal()
	}
	m.mutex.Unlock()

	event := models.NewMetronomeStateChangeEvent(isRunning, m.bpm)
	m.dispatcher.Dispatch(event)
	return nil
}

func (m *MetronomeService) GetState() models.MetronomeStateEventData {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return models.MetronomeStateEventData{
		IsRunning: m.ticker != nil,
		Bpm:       m.bpm,
	}
}

func (m *MetronomeService) startInternal() {
	interval := time.Minute / time.Duration(m.bpm)
	m.ticker = time.NewTicker(interval)
	m.beat()
	go func() {
		for range m.ticker.C {
			m.beat()
		}
	}()
}

func (m *MetronomeService) stopInternal() {
	m.ticker.Stop()
	m.ticker = nil
	m.beatCount = 0
}
