package repositories

import "github.com/dualgiupponi/gestor_simposios_universidad/domain"

type AsistenteRepository interface {
	GetAsistente(id int) (*domain.Asistente, error)
	CreateAsistente(nuevoAsistente domain.Asistente) error
	UpdateAsistente(asistenteActualizado domain.Asistente) error
}
