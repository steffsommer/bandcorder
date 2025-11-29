package services

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
)

type SoundEffectPlayer struct {
	eventBus        interfaces.EventBus
	playbackService interfaces.PlaybackService
}

func NewSoundEffectPlayer(
	eventBus interfaces.EventBus,
	playbackService interfaces.PlaybackService,
) *SoundEffectPlayer {
	return &SoundEffectPlayer{
		eventBus:        eventBus,
		playbackService: playbackService,
	}
}

func (s *SoundEffectPlayer) Init(
	eventToAudioEffect map[models.EventId]interfaces.AudioEffect,
) {
	for eventId, audioEffect := range eventToAudioEffect {
		s.eventBus.OnEvent(eventId, func(_ any) {
			go s.playbackService.Play(audioEffect)
		})
	}
}
