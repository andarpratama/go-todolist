package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB // Exported DB variable

// Initialize MySQL connection
func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_todolist")
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	log.Println("Connected to MySQL database!")
}
