package repository

import (
	server "example/hello"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (auth *AuthRepository) CreateUser(user server.User) (int64, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `username`, `password_hash`) values (?, ?, ?)", "users")

	result, err := auth.db.Exec(query, user.Name, user.UserName, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

func (auth *AuthRepository) GetUserById(id int) (server.User, error) {

}
