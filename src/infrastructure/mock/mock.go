package mock

type Mock struct{}

func New() *Mock {
	return &Mock{}
}
