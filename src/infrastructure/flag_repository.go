package infrastructure

import (
	"database/sql"

	"app/domain/flag"
)

type FlagRepository struct {
	db *sql.DB
}

func NewFlagRepository(db *sql.DB) *FlagRepository {
	return &FlagRepository{
		db: db,
	}
}

func (fr *FlagRepository) Create(f *flag.Flag) error {
	query := "INSERT INTO flags(name) VALUES (?)"
	result, err := fr.db.Exec(query, f.Name)
	if err != nil {
		return err
	}

	f.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (fr *FlagRepository) Update(f *flag.Flag) error {
	query := "UPDATE flags SET name = ? WHERE id = ?"

	_, err := fr.db.Exec(query, f.Name, f.ID)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FlagRepository) Delete(id uint64) error {
	query := "DELETE FROM flags WHERE id = ?"
	_, err := fr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (fr *FlagRepository) FindByID(id uint64) (*flag.Flag, error) {
	query := "SELECT id, name FROM flags WHERE id = ?"
	row := fr.db.QueryRow(query, id)

	f := &flag.Flag{}

	err := row.Scan(&f.ID, &f.Name)
	if err != nil {
		return nil, err
	}

	return f, nil

}

func (fr *FlagRepository) FindAll() ([]*flag.Flag, error) {
	query := "SELECT id, name FROM flags"
	rows, err := fr.db.Query(query)
	if err != nil {
		return nil, err
	}

	flags := make([]*flag.Flag, 0)

	for rows.Next() {
		f := &flag.Flag{}

		err := rows.Scan(&f.ID, &f.Name)
		if err != nil {
			return nil, err
		}

		flags = append(flags, f)
	}

	return flags, nil

}
