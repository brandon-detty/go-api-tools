package repo

// A Repo abstracts CRUD operations for a specific persistence strategy;
// stores items of type T with IDs of type I
type Repo[T any, I comparable] interface {
	Set(id I, item *T) (*T, error)
	Get(id I) (*T, error)
	Delete(id I) error
}
