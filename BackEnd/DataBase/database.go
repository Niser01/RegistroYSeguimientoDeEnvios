package DataBase

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("POSTGRES_USER")
	db_password := os.Getenv("POSTGRES_PASSWORD")
	db_database := os.Getenv("POSTGRES_DB")
	db_port := os.Getenv("POSTGRES_PORT")
	db_sslmode := os.Getenv("SSL_MODE")

	db_connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", db_host, db_user, db_password, db_database, db_port, db_sslmode)

	DB, err = gorm.Open(postgres.Open(db_connection), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return err
	}
	fmt.Printf("Conexión con la Base de Datos establecida correctamente.")
	return nil
}
