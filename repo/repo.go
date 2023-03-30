package repo

import e "github.com/brandon-detty/go-api-tools/entity"

// A Repo abstracts CRUD operations for a specific persistence strategy;
// stores items of type T with IDs of type I
type Repo[T e.Entity[I], I comparable] interface {
	Save(item *T) (*T, error)
	Get(id I) (*T, error)
	Delete(id I) error
}
