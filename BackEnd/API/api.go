package API

import (
	"BackEnd/Controller"

	"github.com/gorilla/mux"
)

func RutasAPI(r *mux.Router) {
	r.HandleFunc("/rastreopedido", Controller.ControllerRastreoPedido).Methods("GET")
	r.HandleFunc("/crearpedido", Controller.ControllerCrearPedido).Methods("POST")
	r.HandleFunc("/crearcliente", Controller.ControllerCrearCliente).Methods("POST")
	r.HandleFunc("/actualizardireccion", Controller.ControllerActualizarDireccion).Methods("PUT")
	r.HandleFunc("/actualizarpedido", Controller.ControllerActualizarEstadoPedido).Methods("PUT")
}
