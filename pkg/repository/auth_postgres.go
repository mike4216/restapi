package repository

import (
	"TODOList_REST"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user TODOList_REST.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING Id", usersTableName)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (TODOList_REST.User, error) {
	var user TODOList_REST.User
	query := fmt.Sprintf("SELECT FROM %s WHERE username=$1 AND password_hash=$2", usersTableName)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
