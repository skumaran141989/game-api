package BusinessLogicLayer

type UpSert[T any] interface {
	Save() (T, error)
}
