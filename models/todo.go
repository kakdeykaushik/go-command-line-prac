package models

// todo: return types to be added
type Todo interface {
	Save()
	MarkComplete()
	Delete()
}

type TodoModel struct {
	Id        string
	Title     string
	Completed bool
}
