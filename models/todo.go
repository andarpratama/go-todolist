package models

import (
	"fmt"
	"go-todolist/config"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoUpdate struct {
	Title     *string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
}

func AddTodo(title string) error { // Exported
	_, err := config.DB.Exec("INSERT INTO todos (title) VALUES (?)", title)
	return err
}

func GetTodos() ([]Todo, error) {
	rows, err := config.DB.Query("SELECT id, title, completed FROM todos")
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

func GetTodoByID(id int) (Todo, error) {
	var todo Todo
	err := config.DB.QueryRow("SELECT id, title, completed FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func PartialUpdateTodo(id string, todoUpdate TodoUpdate) error {
	// Build the update query dynamically based on provided fields
	query := "UPDATE todos SET"
	params := []interface{}{}

	if todoUpdate.Title != nil {
		query += " title = ?,"
		params = append(params, *todoUpdate.Title)
	}
	if todoUpdate.Completed != nil {
		query += " completed = ?,"
		params = append(params, *todoUpdate.Completed)
	}

	// Remove trailing comma and add WHERE clause
	if len(params) == 0 {
		return nil // No fields to update
	}
	query = query[:len(query)-1] + " WHERE id = ?"
	params = append(params, id)

	_, err := config.DB.Exec(query, params...)
	return err
}

func DeleteTodoByID(id string) error {
	// Attempt to delete the todo item by ID
	result, err := config.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	// Check if the deletion affected any rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no todo found with id: %s", id) // Return an error if no rows were deleted
	}

	return nil
}
