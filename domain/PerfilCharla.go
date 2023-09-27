package domain

import "time"

type PerfilCharla struct {
	Id                 int
	Titulo             string
	Descripcion        string
	FechaAlta          time.Time
	FechaActualizacion time.Time
	FechaBaja          *time.Time
}
