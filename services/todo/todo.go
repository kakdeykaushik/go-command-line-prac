package todo

import (
	"fmt"
	"go-cli-p/models"

	"github.com/google/uuid"
)

type Todo struct {
	Id       string
	Title    string
	Complete bool
	storage  models.Storage
}

func NewTodo(title string, storage models.Storage) *Todo {
	id := uuid.New().String()
	return &Todo{
		Id:       id,
		Title:    title,
		Complete: false,
		storage:  storage,
	}
}

// todo: return types to be added
func (t *Todo) Save() {
	t.storage.Save(t.Id, t)
}

func (t *Todo) MarkComplete() {
	t.Complete = true
	t.Save()
}

func (t *Todo) Delete() {
	t.storage.DeleteByID(t.Id)
}

func (t *Todo) UpdateTitle(newTitle string) {
	t.Title = newTitle
	t.Save()
}

func (t *Todo) GetDetails() string {
	status := "Incomplete"
	if t.Complete {
		status = "Complete"
	}
	return fmt.Sprintf("Id: %s, Title: %s, Status: %s", t.Id, t.Title, status)
}

func (t Todo) String() string {
	return fmt.Sprintf("Id: %s Title: %s Complete: %v", t.Id, t.Title, t.Complete)
}
