package conectar

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importar el driver de MySQL
	"github.com/joho/godotenv"
)

var Db *sql.DB

func Conectar() {
	err := godotenv.Load()
	if err != nil {
		panic("Error al cargar las variables de entorno: " + err.Error())
	}

	// Obtener variables de entorno
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST") // e.g. "localhost"
	port := os.Getenv("DB_PORT") // e.g. "3306"
	dbname := os.Getenv("DB_NAME")

	// Construir DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	// Abrir conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Error al abrir la conexión a la base de datos: " + err.Error())
	}

	// Probar la conexión
	if err = db.Ping(); err != nil {
		panic("No se pudo conectar a la base de datos: " + err.Error())
	}

	Db = db // Asignar a la variable global
}
