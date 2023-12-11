package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	conector := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", conector)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createTable(db)

	defer db.Close()
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS task (
    	id SERIAL PRIMARY KEY,
    	name VARCHAR(100) NOT NULL,
    	finished BOOLEAN,
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}
