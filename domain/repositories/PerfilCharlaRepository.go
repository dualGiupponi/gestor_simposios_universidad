package repositories

import "github.com/dualgiupponi/gestor_simposios_universidad/domain"

type PerfilCharlaRepository interface {
	GetPerfilCharla(id int) (*domain.PerfilCharla, error)
	GetAllPerfilCharla() ([]domain.PerfilCharla, error)
	CreatePerfilCharla(nuevoAsistente domain.PerfilCharla) error
	UpdatePerfilCharla(asistenteActualizado domain.PerfilCharla) error
}
