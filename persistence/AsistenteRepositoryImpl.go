package persistence

import (
	"database/sql"

	"github.com/dualgiupponi/gestor_simposios_universidad/domain"
)

type AsistenteRepositoryImplementation struct {
	*sql.DB
}

func NewAsistenteRepositoryImplementation(db *sql.DB) AsistenteRepositoryImplementation {
	return AsistenteRepositoryImplementation{db}
}

func (i *AsistenteRepositoryImplementation) GetAsistente(id int) (*domain.Asistente, error) {
	const consultaSQL = `SELECT 
							asis_id,
							asis_nom,
							asis_ape,
							asis_url_foto,
							asis_email,
							asis_tel,
							asis_din,
							asis_fec_alt,
							asis_fec_act,
							asis_fec_baj,
							pflc_id,
							pflc_tit,
							pflc_des,
							pflc_fec_alt,
							pflc_fec_act,
							pflc_fec_baj
						FROM asistentes
						INNER JOIN perfiles_charla
							ON asis_pflc_id = pflc_id
						WHERE
							asis_id = $1
							AND asis_fec_alt is null`

	rows := i.DB.QueryRow(consultaSQL, id)
	defer rows.Err().Error()

	if rows.Err() == sql.ErrNoRows {
		return nil, nil
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var asistenteBuscado asistenteDb
	var perfilCharla perfilCharlaDB

	rows.Scan(
		&asistenteBuscado.Id,
		&asistenteBuscado.Nombre,
		&asistenteBuscado.Apellido,
		&asistenteBuscado.Url_Foto,
		&asistenteBuscado.Email,
		&asistenteBuscado.Telefono,
		&asistenteBuscado.Dinero,
		&asistenteBuscado.FechaAlta,
		&asistenteBuscado.FechaActualizacion,
		&asistenteBuscado.FechaBaja,
		&perfilCharla.Id,
		&perfilCharla.Titulo,
		&perfilCharla.Descripcion,
		&perfilCharla.FechaAlta,
		&perfilCharla.FechaActualizacion,
		&perfilCharla.FechaBaja,
	)

	asistenteEntidad := convertDbToAsistente(asistenteBuscado, perfilCharla)
	return &asistenteEntidad, nil
}
func (i *AsistenteRepositoryImplementation) CreateAsistente(nuevoAsistente domain.Asistente) error {
	panic("Method not implemented")
}
func (i *AsistenteRepositoryImplementation) UpdateAsistente(asistenteActualizado domain.Asistente) error {
	panic("Method not implemented")
}
