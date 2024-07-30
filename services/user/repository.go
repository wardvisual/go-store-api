package user

import (
	"database/sql"

	"github.com/wardvisual/go-store-api/types"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateUser(user types.User) error {
	var statement string = "INSERT INTO users (firstName, lastName, email, password), VALUES(?,?,?,?)"
	_, err := r.db.Exec(statement, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
