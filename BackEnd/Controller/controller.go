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

func ControllerCrearPedido(w http.ResponseWriter, r *http.Request) {
	var body_input Services.CrearPedidoInput
	err := json.NewDecoder(r.Body).Decode(&body_input)
	if err != nil {
		http.Error(w, "JSON Invalido", http.StatusBadRequest)
		return
	}

	codigoTracking, err := Services.CrearPedido(body_input)
	if err != nil {
		mensaje := fmt.Sprintf("Error al crear el nuevo pedido. - %v", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"codigo_tracking": codigoTracking,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta a JSON", http.StatusInternalServerError)
		return
	}
}

func ControllerCrearCliente(w http.ResponseWriter, r *http.Request) {
	var body_input Services.CrearClienteInput

	err := json.NewDecoder(r.Body).Decode(&body_input)
	if err != nil {
		http.Error(w, "JSON Invalido", http.StatusBadRequest)
		return
	}

	err = Services.CrearCliente(body_input)
	if err != nil {
		mensaje := fmt.Sprintf("Error al crear un nuevo cliente. - %v", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
