package repository

import (
	"github.com/bookmarks-api/models"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) Authorization {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) AddUser(user *models.User) (int, error) {
	query := `insert into users (username, email, password_hash) values($1, $2, $3) returning id`

	var id int
	err := r.db.QueryRow(query, user.UserName, user.Email, user.Password).Scan(&id)

	return id, err
}

func (r *AuthorizationRepository) GetUserId(email, password string) (int, error) {
	query := `select id from users where email=$1 and password_hash=$2`

	var id int
	err := r.db.Get(&id, query, email, password)
	return id, err
}
