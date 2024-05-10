package service

import (
	"crypto/sha1"
	server "example/hello"
	"example/hello/pkg/repository"
	"fmt"
)

const salt = "ddfsdfr24we34sdjh347980-1"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (auth *AuthService) CreateUser(user server.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	userId, err := auth.repo.CreateUser(user)
	auth.repo.GetUserById()

	return userId, err
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
