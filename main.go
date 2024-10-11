package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()

	http.HandleFunc("/todos/create", createTodoHandler)
	http.HandleFunc("/todos", getAllTodosHandler)
	http.HandleFunc("/todo", getTodoByIDHandler)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
