package repository

type BarRepository interface {
	GetAll() ([]string, error)
	Insert(item string) error
	Clear() bool
}

// compile-time check
var _ BarRepository = (*BarRepositoryImpl)(nil)

type BarRepositoryImpl struct {
}

func NewBarRepository() *BarRepositoryImpl {
	return &BarRepositoryImpl{}
}

func (r *BarRepositoryImpl) GetAll() ([]string, error) {
	return []string{"dummy1", "dummy2"}, nil
}

func (r *BarRepositoryImpl) Insert(item string) error {
	return nil
}

func (r *BarRepositoryImpl) Clear() bool {
	return true
}
