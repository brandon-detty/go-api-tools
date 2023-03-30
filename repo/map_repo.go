package repo

import (
	"fmt"

	e "github.com/brandon-detty/go-api-tools/entity"
)

type MapRepo[T e.Entity[I], I comparable] struct {
	store map[I]*T
}

func NewMapRepo[T e.Entity[I], I comparable]() *MapRepo[T, I] {
	store := make(map[I]*T)
	r := MapRepo[T, I]{
		store,
	}
	return &r
}

func (r *MapRepo[T, I]) Save(item *T) (*T, error) {
	id := (*item).Id()
	r.store[id] = item
	return item, nil
}

func (r *MapRepo[T, I]) Get(id I) (*T, error) {
	if item, ok := r.store[id]; ok {
		return item, nil
	}
	return nil, fmt.Errorf("item with id '%v' not found", id)
}

func (r *MapRepo[T, I]) Delete(id I) error {
	if _, ok := r.store[id]; !ok {
		return fmt.Errorf("item with id '%v' not found", id)
	}
	delete(r.store, id)
	return nil
}
