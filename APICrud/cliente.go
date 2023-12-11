package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "tu_usuario"
	dbPassword = "tu_contraseña"
	dbName     = "todolist"
	dbHost     = "localhost"
	dbPort     = "5432"
)

func main() {
	// Cadena de conexión
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	// Conectarse a la base de datos
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

	// Menú principal
	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("1. Mostrar Tareas")
		fmt.Println("2. Crear Nueva Tarea")
		fmt.Println("0. Salir")

		var opcion int
		fmt.Print("Selecciona una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			err := mostrarTareas(db)
			if err != nil {
				log.Println("Error al mostrar las tareas:", err)
			}
		case 2:
			err := crearNuevaTarea(db)
			if err != nil {
				log.Println("Error al crear una nueva tarea:", err)
			}
		case 0:
			fmt.Println("¡Adiós!")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Inténtalo de nuevo.")
		}
	}
}

func mostrarTareas(db *sql.DB) error {
	fmt.Println("\n--- Mostrar Tareas ---")
	return listarTareas(db)
}

func crearNuevaTarea(db *sql.DB) error {
	fmt.Println("\n--- Crear Nueva Tarea ---")

	// Obtener valores desde la entrada estándar
	var nombreTarea string
	var hecho bool

	fmt.Print("Nombre de la tarea: ")
	fmt.Scanln(&nombreTarea)

	fmt.Print("¿La tarea está hecha? (true/false): ")
	fmt.Scanln(&hecho)

	// Llamar a la función de inserción
	err := insertarTarea(db, nombreTarea, hecho)
	if err != nil {
		return err
	}

	fmt.Printf("Tarea '%s' insertada con éxito.\n", nombreTarea)
	return nil
}

func listarTareas(db *sql.DB) error {
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
