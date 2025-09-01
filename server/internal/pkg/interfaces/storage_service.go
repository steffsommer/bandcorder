package interfaces

type StorageService interface {
	Save(fileName string, data []float32) error
}
