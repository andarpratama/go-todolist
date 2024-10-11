// handlers.go
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := addTodo(todo.Title); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Todo created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func getAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := getTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := getTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}