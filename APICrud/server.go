package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "mi_contraseña"
	dbName     = "todolist"
	dbHost     = "localhost" // Cambia esto según tu configuración de Docker
	dbPort     = "5432"      // Cambia esto según tu configuración de Docker
)

func main() {
	// Cadena de conexión
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, "postgres", dbHost, dbPort)

	// Conectarse a la base de datos principal (postgres)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos PostgreSQL!")

	// Verificar si la base de datos existe
	if err := createDatabase(db, dbName); err != nil {
		log.Fatal(err)
	}

	// Cadena de conexión para la base de datos específica (todolist)
	connectionString = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	// Conectarse a la base de datos todolist
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar la conexión a la base de datos todolist
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Conexión exitosa a la base de datos '%s' en PostgreSQL!\n", dbName)

	createTable(db)

	// Llamar a la función de inserción
	// Obtener valores desde la entrada estándar
	var nombreTarea string
	var hecho bool

	fmt.Print("Nombre de la tarea: ")
	fmt.Scanln(&nombreTarea)

	fmt.Print("¿La tarea está hecha? (true/false): ")
	fmt.Scanln(&hecho)

	// Llamar a la función de inserción
	err = insertarTarea(db, nombreTarea, hecho)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tarea '%s' insertada con éxito.\n", nombreTarea)

	mostrarTareas(db)
}

// createDatabase verifica si la base de datos existe y la crea si no.
func createDatabase(db *sql.DB, dbName string) error {
	// Consultar si la base de datos existe
	query := "SELECT 1 FROM pg_database WHERE datname = $1"
	var exists bool
	err := db.QueryRow(query, dbName).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// Si no existe, crear la base de datos
	if !exists {
		createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		_, err := db.Exec(createQuery)
		if err != nil {
			return err
		}
		fmt.Printf("Base de datos '%s' creada con éxito.\n", dbName)
	}

	return nil
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		done boolean
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func insertarTarea(db *sql.DB, nombre string, hecho bool) error {
	query := "INSERT INTO tasks (name, done) VALUES ($1, $2)"
	_, err := db.Exec(query, nombre, hecho)
	return err
}

func mostrarTareas(db *sql.DB) error {
	query := "SELECT * FROM tasks"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Todas las tareas en la tabla 'tasks':")
	fmt.Println("===================================")

	for rows.Next() {
		var id int
		var name string
		var done bool
		err := rows.Scan(&id, &name, &done)
		if err != nil {
			return err
		}
		fmt.Printf("ID: %d, Nombre: %s, Hecho: %t\n", id, name, done)
	}

	return nil
}
