package Services

import (
	"BackEnd/DataBase"
	"BackEnd/Model"
)

type RastreoPedidoSalida struct {
	Estado           string `json:"estado"`
	Descripcion      string `json:"descripcion"`
	DireccionDestino string `json:"direccion_destino"`
}

func RastreoPedido(codigoTracking string) (RastreoPedidoSalida, error) {
	var envio Model.Envio
	var direccion Model.Direccion
	var estado Model.EstadoEnvio
	var salida RastreoPedidoSalida

	// Buscar el envío por codigoTracking
	if err := DataBase.DB.Where("codigo_tracking = ?", codigoTracking).First(&envio).Error; err != nil {
		return salida, err
	}
	// Buscar la dirección destino usando ClienteDestinoID
	if err := DataBase.DB.Where("cliente_id = ?", envio.ClienteDestinoID).First(&direccion).Error; err != nil {
		return salida, err
	}
	// Buscar el último estado (más reciente por fecha_estado)
	if err := DataBase.DB.Where("envio_id = ?", envio.ID).Order("fecha_estado desc").First(&estado).Error; err != nil {
		return salida, err
	}
	// Armar struct de salida
	salida = RastreoPedidoSalida{
		Estado:           estado.Estado,
		Descripcion:      estado.DescripcionEnvio,
		DireccionDestino: direccion.Direccion,
	}

	return salida, nil
}
