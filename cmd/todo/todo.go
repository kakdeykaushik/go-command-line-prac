package todo

import (
	"errors"
	"go-cli-p/logger"
	"go-cli-p/models"
	"go-cli-p/services/todo"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	store        models.Storage = models.NewStorage("memory")
	setFlag      string
	completeFlag string
	updateFlag   string
	deleteFlag   string
	viewFlag     string
)

var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "*le me again TODO",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if !validFlags() {
			logger.Fatal("at least one flag must be provided")
		}

		if setFlag != "" {
			todoTitle := setFlag
			newTodo := todo.NewTodo(todoTitle, store)
			newTodo.Save()
			logger.Printf("%+v\n", newTodo)
		}

		if completeFlag != "" {
			todoId := completeFlag
			theTodo, _ := store.GetByID(todoId).(*todo.Todo)
			if theTodo == nil {
				logger.Println("Todo does not exists")
				return
			}
			theTodo.MarkComplete()
		}

		if updateFlag != "" {
			updateString := updateFlag
			todoId, title, err := splitString(updateString)
			if err != nil {
				logger.Println(err)
				return
			}

			theTodo, _ := store.GetByID(todoId).(*todo.Todo)
			if theTodo == nil {
				logger.Println("Todo does not exists")
				return
			}

			theTodo.UpdateTitle(title)
		}

		if deleteFlag != "" {
			todoId := deleteFlag
			theTodo, _ := store.GetByID(todoId).(*todo.Todo)
			if theTodo == nil {
				logger.Println("Todo does not exists")
				return
			}
			theTodo.Delete()
		}

		if viewFlag != "" {
			todoId := viewFlag
			todo := todo.GetTaskDetails(todoId)
			if todo == nil {
				logger.Println("Todo does not exists")
				return
			}
			logger.Println("title : " + todo.Title)
			logger.Println("Completed: " + strconv.FormatBool(todo.Completed))
		}
		// logger.Println(store.ViewAll())
	},
}

func init() {
	TodoCmd.Flags().StringVarP(&setFlag, "set", "s", "", "save new todo")
	TodoCmd.Flags().StringVarP(&completeFlag, "complete", "c", "", "mark todo as complete")
	TodoCmd.Flags().StringVarP(&updateFlag, "update", "u", "", "updates todo")
	TodoCmd.Flags().StringVarP(&deleteFlag, "delete", "d", "", "deletes todo")
	TodoCmd.Flags().StringVarP(&viewFlag, "view", "v", "", "details of todo")
}

// helper
func splitString(input string) (string, string, error) {
	parts := strings.Split(input, " ")
	if len(parts) >= 2 {
		firstPart := parts[0]
		secondPart := strings.Join(parts[1:], " ")

		return firstPart, secondPart, nil
	}

	return "", "", errors.New("unsplitable string")
}

// helper
func validFlags() bool {
	return !(setFlag == "" && completeFlag == "" && updateFlag == "" && deleteFlag == "" && viewFlag == "")
}
