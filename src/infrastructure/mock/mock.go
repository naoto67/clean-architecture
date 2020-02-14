package mock

type MemoryMock struct{}

func NewMemoryMock() *MemoryMock {
	return &MemoryMock{}
}
