package Services

import (
	"BackEnd/DataBase"
	"BackEnd/Model"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
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

type CrearPedidoInput struct {
	CorreoClienteOrigen  string    `json:"correo_cliente_origen"`
	CorreoClienteDestino string    `json:"correo_cliente_destino"`
	Descripcion          string    `json:"descripcion"`
	Peso                 float64   `json:"peso"`
	ValorEnvio           float64   `json:"valor_envio"`
	Moneda               string    `json:"moneda"`
	FechaEntregaEstimada time.Time `json:"fecha_entrega_estimada"`
}

func generarCodigoTrackingUnico(db *gorm.DB) (string, error) {
	var codigo string
	for {
		codigo = fmt.Sprintf("TRK-%d", time.Now().UnixNano())

		var count int64
		if err := db.Model(&Model.Envio{}).Where("codigo_tracking = ?", codigo).Count(&count).Error; err != nil {
			return "", fmt.Errorf("error verificando codigo_tracking: %w", err)
		}

		if count == 0 {
			break
		}
	}
	return codigo, nil
}
func CrearPedido(input CrearPedidoInput) (string, error) {
	// Buscar cliente origen por correo
	var clienteOrigen, clienteDestino Model.Cliente
	if err := DataBase.DB.Where("correo = ?", input.CorreoClienteOrigen).First(&clienteOrigen).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("cliente origen con correo %s no encontrado", input.CorreoClienteOrigen)
		}
		return "", fmt.Errorf("error consultando cliente origen: %w", err)
	}

	// Buscar cliente destino por correo
	if err := DataBase.DB.Where("correo = ?", input.CorreoClienteDestino).First(&clienteDestino).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("cliente destino con correo %s no encontrado", input.CorreoClienteDestino)
		}
		return "", fmt.Errorf("error consultando cliente destino: %w", err)
	}

	// Generar código tracking único
	codigoTracking, err := generarCodigoTrackingUnico(DataBase.DB)
	if err != nil {
		return "", err
	}

	// Crear el envío
	envio := Model.Envio{
		CodigoTracking:       codigoTracking,
		ClienteOrigenID:      clienteOrigen.ID,
		ClienteDestinoID:     clienteDestino.ID,
		Descripcion:          input.Descripcion,
		Peso:                 input.Peso,
		ValorEnvio:           input.ValorEnvio,
		Moneda:               input.Moneda,
		FechaEntregaEstimada: input.FechaEntregaEstimada,
		Entregado:            false,
	}

	if err := DataBase.DB.Create(&envio).Error; err != nil {
		return "", fmt.Errorf("error creando el pedido: %w", err)
	}

	// Crear estado inicial
	estado := Model.EstadoEnvio{
		EnvioID:          envio.ID,
		Estado:           "El envío ha sido recibido en la bodega de Pack&Track",
		DescripcionEnvio: "El pedido ha sido registrado y está en bodega de Pack&Track.",
	}

	if err := DataBase.DB.Create(&estado).Error; err != nil {
		return "", fmt.Errorf("error creando estado inicial: %w", err)
	}

	return codigoTracking, nil
}

type CrearClienteInput struct {
	Nombre       string `json:"nombre"`
	Correo       string `json:"correo"`
	Telefono     string `json:"telefono"`
	Direccion    string `json:"direccion"`
	Ciudad       string `json:"ciudad"`
	Departamento string `json:"departamento"`
	Pais         string `json:"pais"`
	CodigoPostal string `json:"codigo_postal"`
}

func CrearCliente(input CrearClienteInput) error {
	// Verificar si ya existe un cliente con ese correo
	var existente Model.Cliente
	if err := DataBase.DB.Where("correo = ?", input.Correo).First(&existente).Error; err == nil {
		// Si no hay error y encontró un registro, significa que ya existe
		return fmt.Errorf("ya existe un cliente con ese correo")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si hubo otro error de base de datos, lo retornamos
		return fmt.Errorf("error al buscar cliente: %w", err)
	}

	// Crear cliente nuevo
	cliente := Model.Cliente{
		Nombre:   input.Nombre,
		Correo:   input.Correo,
		Telefono: input.Telefono,
	}

	if err := DataBase.DB.Create(&cliente).Error; err != nil {
		return fmt.Errorf("error al crear cliente: %w", err)
	}

	// Crear dirección asociada
	direccion := Model.Direccion{
		ClienteID:    cliente.ID,
		Direccion:    input.Direccion,
		Ciudad:       input.Ciudad,
		Departamento: input.Departamento,
		Pais:         input.Pais,
		CodigoPostal: input.CodigoPostal,
	}

	if err := DataBase.DB.Create(&direccion).Error; err != nil {
		return fmt.Errorf("error al crear dirección: %w", err)
	}

	return nil
}
