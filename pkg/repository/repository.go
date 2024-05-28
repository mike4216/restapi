package repository

import (
	"TODOList_REST"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user TODOList_REST.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
