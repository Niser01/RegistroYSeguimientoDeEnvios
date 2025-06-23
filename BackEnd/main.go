package main

import (
	"BackEnd/API"
	"BackEnd/DataBase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := DataBase.InitDB()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	log.Println("Conexión con la base de datos establecida correctamente.")

	route_handler := mux.NewRouter()
	API.RutasAPI(route_handler)

	handler := enableCORS(route_handler)
	log.Println("Backend inicializado correctamente")

	//Se lanza el servidor en el puerto 8080
	log.Println("Servidor escuchando en el puerto 8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Configuración CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Manejo de preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pasar al siguiente handler
		h.ServeHTTP(w, r)
	})
}
