package main

import (
	"go-todolist/config"
	"go-todolist/handlers"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Define route handlers
	http.HandleFunc("/todos/create", handlers.CreateTodoHandler) // Create a new todo
	http.HandleFunc("/todos", handlers.GetAllTodosHandler)       // Get all todos
	http.HandleFunc("/todo", handlers.GetTodoByIDHandler)        // Get a todo by ID
	http.HandleFunc("/todos/update", handlers.UpdateTodoHandler) // Update a todo
	http.HandleFunc("/todos/delete", handlers.DeleteTodoHandler) // Delete a todo

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
