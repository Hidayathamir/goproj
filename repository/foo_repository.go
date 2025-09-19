package repository

type FooRepository interface {
	FindByID(id int) string
	Save(data string) error
	Delete(id int) bool
}

// compile-time check
var _ FooRepository = (*FooRepositoryImpl)(nil)

type FooRepositoryImpl struct {
}

func NewFooRepository() *FooRepositoryImpl {
	return &FooRepositoryImpl{}
}

func (r *FooRepositoryImpl) FindByID(id int) string {
	return "dummy result"
}

func (r *FooRepositoryImpl) Save(data string) error {
	return nil
}

func (r *FooRepositoryImpl) Delete(id int) bool {
	return true
}
