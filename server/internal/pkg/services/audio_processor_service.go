package services

import (
	"github.com/mjibson/go-dsp/fft"
	"math"
	"math/cmplx"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"server/internal/pkg/utils"
)

const (
	FFTSize  = 2048   // Number of samples to accumulate before data is processed
	BarCount = 40     // Number of frequency bar data (range 0-100 per bar) to calculate
	LowFreq  = 50.0   // approximate frequency of the lowest bar
	HighFreq = 6000.0 // approximate frequency of the highest bar
)

type AudioSampleProcessorService struct {
	dispatcher   interfaces.EventDispatcher
	sampleBuffer []float32
}

// NewAudioProcessorService creates a new AudioProcessorService
func NewAudioProcessorService(
	dispatcher interfaces.EventDispatcher,
) *AudioSampleProcessorService {
	return &AudioSampleProcessorService{
		dispatcher:   dispatcher,
		sampleBuffer: make([]float32, 0, FFTSize),
	}
}

// Process buffers audio samples and dispatches a LiveAudioDataEvent when
// FFTSize samples accumulated over a internal threshold and got processed.
// It calculates loudness and frequency bars, sending empty bars if the audio
// is too quiet (loudness < 2). The sample buffer uses a 50% overlap window for
// smoother transitions between frames.
func (a *AudioSampleProcessorService) Process(samples []float32) {
	a.sampleBuffer = append(a.sampleBuffer, samples...)
	if len(a.sampleBuffer) < FFTSize {
		return
	}

	loudness := calculateRootMeanSqaureLoudness(a.sampleBuffer[:FFTSize])

	// Remove noise
	var bars []int
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

func calculateRootMeanSqaureLoudness(samples []float32) uint8 {
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

	lowBin := (LowFreq / nyquist) * float64(FFTSize/2)
	highBin := (HighFreq / nyquist) * float64(FFTSize/2)

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

