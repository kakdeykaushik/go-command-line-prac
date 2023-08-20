package models

// todo: return types to be added
type Todo interface {
	Save()
	MarkComplete()
	Delete()
}
