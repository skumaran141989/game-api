package BusinessLogicLayer

type Extract[T any] interface {
	Get() (T, error)
}
