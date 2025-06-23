package Controller

import (
	"BackEnd/Services"
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func ControllerRastreoPedido(w http.ResponseWriter, r *http.Request) {
	codigoTracking := r.URL.Query().Get("codigoTracking")
	if codigoTracking == "" {
		writeJSONError(w, http.StatusBadRequest, "La variable 'codigoTracking' no se encuentra entre los parametros")
		return
	}
	query, err := Services.RastreoPedido(codigoTracking)
	if err != nil {
		mensaje := fmt.Sprintf("Error al buscar el pedido: %v - %v", codigoTracking, err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta a JSON", http.StatusInternalServerError)
		return
	}
}
