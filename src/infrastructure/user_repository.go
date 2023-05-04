package infrastructure

import (
	"database/sql"

	"projeto/app/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(u *user.User) error {
	query := "INSERT INTO users(name, email, password, profile) VALUES (?, ?, ?, ?)"
	result, err := ur.db.Exec(query, u.Name, u.Email, u.Password, u.Profile)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Update(u *user.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ?, profile = ? WHERE id = ?"
	_, err := ur.db.Exec(query, u.Name, u.Email, u.Password, u.Profile, u.ID)
	if err != nil {
		return err
	}

	return nil
	_, err := ur.db.Exec(query, u.Name, u.Email, u.Password, u.Profile, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(id uint64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := ur.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) FindByID(id uint64) (*user.User, error) {
	query := "SELECT id, name, email, password, profile FROM users WHERE id = ?"
	row := ur.db.QueryRow(query, id)

	u := &user.User{}

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Profile)
	if err != nil {
		return nil, err
	}

	return u, nil

}

func (ur *UserRepository) FindByEmail(email string) (*user.User, error) {
	query := "SELECT id, name, email, password, profile FROM users WHERE email = ?"
	row := ur.db.QueryRow(query, email)

	u := &user.User{}

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Profile)
	if err != nil {
		return nil, err
	}

	return u, nil

}

func (ur *UserRepository) FindAll() ([]*user.User, error) {
	query := "SELECT id, name, email, password, profile FROM users"
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := make([]*user.User, 0)

	for rows.Next() {
		u := &user.User{}

		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Profile)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil

}
