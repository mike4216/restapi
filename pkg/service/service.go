package service

import (
	"TODOList_REST"
	"TODOList_REST/pkg/repository"
)

type Authorization interface {
	CreateUser(user TODOList_REST.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
