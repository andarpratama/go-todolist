package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize MySQL connection
func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_todolist")
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	log.Println("Connected to MySQL database!")
}

func addTodo(title string) error {
	_, err := db.Exec("INSERT INTO todos (title) VALUES (?)", title)
	return err
}

func getTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func getTodoByID(id int) (Todo, error) {
	var todo Todo
	err := db.QueryRow("SELECT id, title, completed FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return todo, err
	}
	return todo, nil
}
