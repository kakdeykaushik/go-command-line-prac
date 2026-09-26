package database

import (
	"database/sql"

	"go-cli-p/logger"
	"go-cli-p/models"
)

// should handel the err

var db *sql.DB

func CreateDatabaseConnection() {
	logger.Println("creating a database connection")
	conn, err := GetDatabaseConnection()
	if err != nil {
		logger.Println("data base creation failed", err)
	}
	db = conn
}

func CreateDatabase() {

	// Create the "todo" table if it doesn't exist
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS todo (
		id TEXT PRIMARY KEY,
		title TEXT,
		completed BOOLEAN
		);`

	_, e := db.Exec(createTableSQL)
	if e != nil {
		logger.Fatal("creation of database failed")
	}

}

func AddTask(taskId string, title string, completed bool) {
	stmt, e := db.Prepare("Insert Into todo Values (?,?,?)")
	if e != nil {
		logger.Fatal("Task creation failed")
	}
	_, err := stmt.Exec(taskId, title, completed)
	if err != nil {
		logger.Println("Task creation failed with error " + err.Error())
	}
}

func GetTask(taskId string) (task models.TodoModel) {
	rows, err := db.Query("Select * from todo where id = ?", taskId)
	if err != nil {
		logger.Fatal("Something went wrong")
	}
	defer rows.Close()
	listedTask := models.TodoModel{}
	for rows.Next() {
		err := rows.Scan(&listedTask.Id, &listedTask.Title, &listedTask.Completed)
		if err != nil {
			logger.Fatal("something went wrong")
		}
	}
	return listedTask
}
