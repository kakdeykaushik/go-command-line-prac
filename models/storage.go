package models

import "go-cli-p/storage"

type Storage interface {
	Save(K, V any)
	GetByID(K any) any
	DeleteByID(K any)
	ViewAll() map[any]any
}

// factory lol
func NewStorage(store string) Storage {
	if store == "memory" {
		return &storage.InMemoryStorage{Items: make(map[any]any)}
	}
	return nil
}
