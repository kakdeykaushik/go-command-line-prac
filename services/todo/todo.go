package todo

import (
	"fmt"
	"go-cli-p/database"
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
	database.AddTask(t.Id, t.Title, t.Complete)
	//t.storage.Save(t.Id, t)
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

// DRY - for future Kaushik
func (t Todo) GetDetails() string {
	status := getStatus(t.Complete)
	return fmt.Sprintf("Id: %s, Title: %s, Status: %s", t.Id, t.Title, status)
}

func (t Todo) String() string {
	status := getStatus(t.Complete)
	return fmt.Sprintf("Id: %s, Title: %s, Status: %s", t.Id, t.Title, status)
}

func getStatus(complete bool) string {
	status := "Incomplete"
	if complete {
		status = "Complete"
	}
	return status
}

func GetTaskDetails(taskId string) *models.TodoModel {
	result := database.GetTask(taskId)
	return &result
}
