package storage

type InMemoryStorage struct {
	// Items should be private, but making this private won;t let factory method(storage.NewStorage) work
	Items map[any]any
}

func (s *InMemoryStorage) Save(K, V any) {
	s.Items[K] = V
}

func (s *InMemoryStorage) GetByID(K any) any {
	return s.Items[K]
}

func (s *InMemoryStorage) DeleteByID(K any) {
	delete(s.Items, K)
}

func (s *InMemoryStorage) ViewAll() map[any]any {
	return s.Items
}
