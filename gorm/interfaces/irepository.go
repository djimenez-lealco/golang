package interfaces

type IRepository[T any, I any] interface {
	Get(id I) *T
	GetAll() *[]T
	Save(*T) *T
	Update(*T) bool
	Delete(id I) bool
}
