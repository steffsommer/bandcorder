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
	dispatcher   interfaces.EventDispatcher
	sampleBuffer []float32
}

func NewAudioProcessorService(
	dispatcher interfaces.EventDispatcher,
) *AudioSampleProcessorService {
	return &AudioSampleProcessorService{
		dispatcher:   dispatcher,
		sampleBuffer: make([]float32, 0, FFTSize),
	}
}

func (a *AudioSampleProcessorService) Process(samples []float32) {
	a.sampleBuffer = append(a.sampleBuffer, samples...)

	// Only process when we have enough samples
	if len(a.sampleBuffer) < FFTSize {
		return
	}

	// Process the first FFTSize samples
	loudness := calculateRMSLoudness(a.sampleBuffer[:FFTSize])

	var bars []int
	// If too quiet, send empty bars
	if loudness < 2 {
		bars = make([]int, BarCount)
	} else {
		bars = calculateFrequencyBars(a.sampleBuffer[:FFTSize])
	}

	event := models.NewLiveAudioDataEvent(loudness, bars)
	a.dispatcher.Dispatch(event)

	// Slide window by half for overlap
	a.sampleBuffer = a.sampleBuffer[FFTSize/2:]
}

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

	spectrum := make([]complex128, FFTSize)
	for i := 0; i < len(audioFrame) && i < FFTSize; i++ {
		spectrum[i] = complex(float64(audioFrame[i]), 0)
	}

	fftResult := fft.FFT(spectrum)

	magnitudes := make([]float64, FFTSize/2)
	for i := 0; i < FFTSize/2; i++ {
		magnitudes[i] = cmplx.Abs(fftResult[i])
	}

	lowFreq := 50.0
	highFreq := 8000.0
	lowBin := (lowFreq / nyquist) * float64(FFTSize/2)
	highBin := (highFreq / nyquist) * float64(FFTSize/2)

	bars := make([]float32, BarCount)
	for i := 0; i < BarCount; i++ {
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
