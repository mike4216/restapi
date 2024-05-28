package service

import (
	"TODOList_REST"
	"TODOList_REST/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "dkljfjsdkjfskldjfklsdjflkjsdklfjsdjf"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user TODOList_REST.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
