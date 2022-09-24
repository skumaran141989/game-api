package repos

type DB[T any] interface {
	Get(id string) (*T, error)
	GetAll(id string) ([]T, error)
	Update(t T) error
	Delete(id string) error
}
