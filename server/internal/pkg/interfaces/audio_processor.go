package interfaces

type AudioProcessor interface {
	Process(audioSample []float32)
}
