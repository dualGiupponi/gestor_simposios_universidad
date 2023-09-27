package domain

import "time"

type Asistente struct {
	Id                 int
	Nombre             string
	Apellido           string
	Url_Foto           string
	Email              string
	Telefono           string
	PerfilInteres      PerfilCharla
	Dinero             float32
	FechaAlta          time.Time
	FechaActualizacion time.Time
	FechaBaja          *time.Time
}
