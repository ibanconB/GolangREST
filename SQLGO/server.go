package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

)

func main() {
	strConn := "user=miuser password=secreto dbname=testDB sslmode=disable"

	db, err := sql.Open("postgres", strConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")

	createTable(db)
	insertTest(db)
}


func createTable(db *sql.DB){
	query:=`CREATE TABLE IF NOT EXISTS prueba(
		id SERIAL PRIMARY KEY,
		name VARCHAR(10) NOT NULL
	)`

	_,err:=db.Exec(query)
	if err!= nil {
		log.Fatal(err)
	}
}

func insertTest(db * sql.DB){
	query:= `INSERT INTO prueba (name) VALUES
	('Nombre1'),
	('Nombre2'),
	('Nombre3'),
	('Nombre4');`

	_,err:= db.Exec(query)
	if err!=nil{
		log.Fatal(err)
	}

	querySelect:= `SELECT * FROM prueba;`

	_,err2:= db.Exec(querySelect)
	if err!=nil{
		log.Fatal(err2)
	}
}