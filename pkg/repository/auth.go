package repository

import (
	"github.com/bookmarks-api/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) Authorization {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) AddUser(user *models.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "start db")
	}

	query := `insert into users (username, email, password_hash) values($1, $2, $3) returning id`

	var id int
	err = r.db.QueryRow(query, user.UserName, user.Email, user.Password).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, "scan items for db")
	}

	return id, nil
}

func (r *AuthorizationRepository) GetUserId(email, password string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "start db")
	}

	query := `select id from users where email=$1 and password_hash=$2`

	var id int
	err = r.db.Get(&id, query, email, password)
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, "get items from db")
	}

	return id, nil
}
