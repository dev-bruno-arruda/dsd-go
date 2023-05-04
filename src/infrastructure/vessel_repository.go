package infrastructure

import (
	"database/sql"
	"projeto/app/domain/vessel"
)

type VesselRepository struct {
	db *sql.DB
}

func NewVesselRepository(db *sql.DB) *VesselRepository {
	return &VesselRepository{
		db: db,
	}
}

func (vr *VesselRepository) Create(v *vessel.Vessel) error {
	query := "INSERT INTO vessels(name, idflag) VALUES (?, ?)"
	result, err := vr.db.Exec(query, v.Name, v.FlagID)
	if err != nil {
		return err
	}

	v.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil

}

func (vr *VesselRepository) Update(v *vessel.Vessel) error {
	query := "UPDATE vessels SET name = ?, idflag = ? WHERE id = ?"
	_, err := vr.db.Exec(query, v.Name, v.FlagID, v.ID)
	if err != nil {
		return err
	}

	return nil

}

func (vr *VesselRepository) Delete(id uint64) error {
	query := "DELETE FROM vessels WHERE id = ?"
	_, err := vr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (vr *VesselRepository) FindByID(id uint64) (*vessel.Vessel, error) {
	query := "SELECT id, name, idflag FROM vessels WHERE id = ?"
	row := vr.db.QueryRow(query, id)

	v := &vessel.Vessel{}

	err := row.Scan(&v.ID, &v.Name, &v.FlagID)
	if err != nil {
		return nil, err
	}

	return v, nil

}

func (vr *VesselRepository) FindAll() ([]*vessel.Vessel, error) {
	query := "SELECT id, name, idflag FROM vessels"
	rows, err := vr.db.Query(query)
	if err != nil {
		return nil, err
	}
	vessels := make([]*vessel.Vessel, 0)

	for rows.Next() {
		v := &vessel.Vessel{}

		err := rows.Scan(&v.ID, &v.Name, &v.FlagID)
		if err != nil {
			return nil, err
		}

		vessels = append(vessels, v)
	}

	return vessels, nil
}
