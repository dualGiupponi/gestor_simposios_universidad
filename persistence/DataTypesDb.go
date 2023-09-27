package persistence

import (
	"database/sql"

	"github.com/dualgiupponi/gestor_simposios_universidad/domain"
)

type asistenteDb struct {
	Id                 int
	Nombre             string
	Apellido           string
	Url_Foto           string
	Email              string
	Telefono           string
	PerfilInteres      int
	Dinero             float32
	FechaAlta          sql.NullTime
	FechaActualizacion sql.NullTime
	FechaBaja          sql.NullTime
}

type perfilCharlaDB struct {
	Id                 int
	Titulo             string
	Descripcion        string
	FechaAlta          sql.NullTime
	FechaActualizacion sql.NullTime
	FechaBaja          sql.NullTime
}

func convertAsistenteToDb(asistente domain.Asistente) asistenteDb {
	asistenteConvertido := asistenteDb{
		Id:            asistente.Id,
		Nombre:        asistente.Nombre,
		Apellido:      asistente.Apellido,
		Url_Foto:      asistente.Url_Foto,
		Email:         asistente.Email,
		Telefono:      asistente.Telefono,
		PerfilInteres: asistente.PerfilInteres.Id,
		Dinero:        asistente.Dinero,
		FechaAlta: sql.NullTime{
			Valid: true,
			Time:  asistente.FechaAlta,
		},
		FechaActualizacion: sql.NullTime{
			Valid: true,
			Time:  asistente.FechaActualizacion,
		},
	}

	if !asistente.FechaBaja.IsZero() {
		asistenteConvertido.FechaBaja = sql.NullTime{
			Valid: true,
			Time:  *asistente.FechaBaja,
		}
	}

	return asistenteConvertido
}

func convertPerfilCharlaToDb(perfil domain.PerfilCharla) perfilCharlaDB {
	perfilConvertido := perfilCharlaDB{
		Id:          perfil.Id,
		Titulo:      perfil.Titulo,
		Descripcion: perfil.Descripcion,
		FechaAlta: sql.NullTime{
			Valid: true,
			Time:  perfil.FechaAlta,
		},
		FechaActualizacion: sql.NullTime{
			Valid: true,
			Time:  perfil.FechaActualizacion,
		},
	}

	if !perfil.FechaBaja.IsZero() {
		perfilConvertido.FechaBaja = sql.NullTime{
			Valid: true,
			Time:  *perfil.FechaBaja,
		}
	}

	return perfilConvertido
}

func convertDbToAsistente(asistente asistenteDb, perfil perfilCharlaDB) domain.Asistente {
	asistenteEntidad := domain.Asistente{
		Id:       asistente.Id,
		Nombre:   asistente.Nombre,
		Apellido: asistente.Apellido,
		Url_Foto: asistente.Url_Foto,
		Email:    asistente.Email,
		Telefono: asistente.Telefono,
		PerfilInteres: domain.PerfilCharla{
			Id:                 perfil.Id,
			Titulo:             perfil.Titulo,
			Descripcion:        perfil.Descripcion,
			FechaAlta:          perfil.FechaAlta.Time,
			FechaActualizacion: perfil.FechaActualizacion.Time,
		},
		Dinero:             asistente.Dinero,
		FechaAlta:          asistente.FechaAlta.Time,
		FechaActualizacion: asistente.FechaActualizacion.Time,
	}

	if asistente.FechaBaja.Valid {
		asistenteEntidad.FechaBaja = &asistente.FechaBaja.Time
	}

	if perfil.FechaBaja.Valid {
		asistenteEntidad.PerfilInteres.FechaBaja = &perfil.FechaBaja.Time
	}

	return asistenteEntidad
}
