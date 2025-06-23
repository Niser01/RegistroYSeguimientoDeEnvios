package Model

import "time"

type Cliente struct {
	ID        int       `gorm:"primaryKey; column:id" json:"id"`
	Nombre    string    `gorm:"column:nombre" json:"nombre"`
	Correo    string    `gorm:"column:correo" json:"correo"`
	Telefono  string    `gorm:"column:telefono" json:"telefono"`
	Creado_en time.Time `gorm:"column:creado_en;autoCreateTime" json:"creado_en"`
}

type Direccion struct {
	ID           int       `gorm:"primaryKey; column:id" json:"id"`
	ClienteID    int       `gorm:"column:cliente_id" json:"cliente_id"`
	Direccion    string    `gorm:"column:direccion" json:"direccion"`
	Ciudad       string    `gorm:"column:ciudad" json:"ciudad"`
	Departamento string    `gorm:"column:departamento" json:"departamento"`
	Pais         string    `gorm:"column:pais" json:"pais"`
	CodigoPostal string    `gorm:"column:codigo_postal" json:"codigo_postal"`
	Creado_en    time.Time `gorm:"column:creado_en;autoCreateTime" json:"creado_en"`
}

type Envio struct {
	ID                   int       `gorm:"primaryKey; column:id" json:"id"`
	CodigoTracking       string    `gorm:"unique; column:codigo_tracking" json:"codigo_tracking"`
	ClienteOrigenID      int       `gorm:"column:cliente_origen_id" json:"cliente_origen_id"`
	ClienteDestinoID     int       `gorm:"column:cliente_destino_id" json:"cliente_destino_id"`
	Descripcion          string    `gorm:"column:descripcion" json:"descripcion"`
	Peso                 float64   `gorm:"column:peso" json:"peso"`
	ValorEnvio           float64   `gorm:"column:valor_envio" json:"valor_envio"`
	Moneda               string    `gorm:"column:moneda" json:"moneda"`
	FechaEnvio           time.Time `gorm:"column:fecha_envio;autoCreateTime" json:"fecha_envio"`
	FechaEntregaEstimada time.Time `gorm:"column:fecha_entrega_estimada" json:"fecha_entrega_estimada"`
	Entregado            bool      `gorm:"column:entregado;default:false" json:"entregado"`
}

type EstadoEnvio struct {
	ID               int       `gorm:"primaryKey; column:id" json:"id"`
	EnvioID          int       `gorm:"column:envio_id" json:"envio_id"`
	Estado           string    `gorm:"column:estado" json:"estado"`
	DescripcionEnvio string    `gorm:"column:descripcion_envio" json:"descripcion_envio"`
	FechaEstado      time.Time `gorm:"column:fecha_estado;autoCreateTime" json:"fecha_estado"`
}
