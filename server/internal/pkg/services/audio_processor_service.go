package services

import (
	"fmt"
	"math"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
)

type AudioSampleProcessorService struct {
	dispatcher interfaces.EventDispatcher
}

func NewAudioProcessorService(
	dispatcher interfaces.EventDispatcher,
) *AudioSampleProcessorService {
	return &AudioSampleProcessorService{
		dispatcher: dispatcher,
	}
}

func (a *AudioSampleProcessorService) Process(samples []float32) {
	loudness := calculateRMSLoudness(samples)
	event := models.NewLiveAudioDataEvent(loudness)
	a.dispatcher.Dispatch(event)
}

// calculateRMSLoudness calculates the loudness of audio samples using RMS (Root Mean Square)
//
// RMS provides a better representation of perceived loudness compared to peak amplitude
// detection by measuring the average energy of a batch of samples.
func calculateRMSLoudness(samples []float32) uint8 {
	fmt.Printf("%v\n", samples)
	if len(samples) == 0 {
		return 0
	}

	var squareSum float64
	for _, sample := range samples {
		squareSum += float64(sample * sample)
	}

	rms := math.Sqrt(squareSum / float64(len(samples)))
	loudness := rms * 100

	if loudness > 100 {
		loudness = 100
	}

	return uint8(math.Round(loudness))
}
