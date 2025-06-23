package main

import (
	"BackEnd/DataBase"
	"log"
	"net/http"
)

func main() {
	err := DataBase.InitDB()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	log.Println("Conexión con la base de datos establecida correctamente.")

	//Se lanza el servidor en el puerto 8080
	log.Println("Servidor escuchando en el puerto 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
