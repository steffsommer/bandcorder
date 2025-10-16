package services

import (
	"math"
	"math/cmplx"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"server/internal/pkg/utils"

	"github.com/mjibson/go-dsp/fft"
)

const (
	FFTSize  = 2048
	BarCount = 40
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

// Process calculates the average loudness of the given audio sample
// and dispatches a LiveAudioDataEvent with the result
func (a *AudioSampleProcessorService) Process(samples []float32) {
	loudness := calculateRMSLoudness(samples)
	bars := calculateFrequencyBars(samples)
	event := models.NewLiveAudioDataEvent(loudness, bars)
	a.dispatcher.Dispatch(event)
}

// calculateRMSLoudness calculates the loudness of audio samples using RMS (Root Mean Square)
//
// RMS provides a better representation of perceived loudness compared to peak amplitude
// detection by measuring the average energy of a batch of samples.
func calculateRMSLoudness(samples []float32) uint8 {
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

func calculateFrequencyBars(audioFrame []float32) []int {
	nyquist := float64(utils.SampleRate) / 2
	// Convert float32 to complex128 for FFT
	spectrum := make([]complex128, FFTSize)
	for i := 0; i < len(audioFrame) && i < FFTSize; i++ {
		spectrum[i] = complex(float64(audioFrame[i]), 0)
	}

	fftResult := fft.FFT(spectrum)

	magnitudes := make([]float64, FFTSize/2)
	for i := 0; i < FFTSize/2; i++ {
		magnitudes[i] = cmplx.Abs(fftResult[i])
	}

	// Map to 40 bars for 50Hz - 8kHz range
	lowFreq := 50.0
	highFreq := 8000.0
	lowBin := (lowFreq / nyquist) * float64(FFTSize/2)
	highBin := (highFreq / nyquist) * float64(FFTSize/2)

	bars := make([]float32, BarCount)
	for i := 0; i < BarCount; i++ {
		// Linear mapping across bin range
		startBin := lowBin + (float64(i)/float64(BarCount))*(highBin-lowBin)
		endBin := lowBin + (float64(i+1)/float64(BarCount))*(highBin-lowBin)

		var sum float64
		count := 0
		for j := int(math.Floor(startBin)); j < int(math.Ceil(endBin)); j++ {
			if j < len(magnitudes) {
				sum += magnitudes[j]
				count++
			}
		}

		if count > 0 {
			bars[i] = float32(sum / float64(count))
		}
	}

	maxMagnitude := float32(0)
	for _, bar := range bars {
		if bar > maxMagnitude {
			maxMagnitude = bar
		}
	}

	if maxMagnitude > 0 {
		for i := range bars {
			bars[i] = (bars[i] / maxMagnitude) * 100
		}
	}

	intBars := make([]int, len(bars))
	for i, v := range bars {
		intBars[i] = int(v)
	}

	return intBars
}
