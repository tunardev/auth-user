package repository

import (
	"database/sql"

	"github.com/tunardev/auth-user/pkg/entity"
)

type UserRepository interface {
	Insert(user entity.User) (int64, error)
	GetByEmail(email string) (entity.User, error)
	GetById(id string) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return userRepository{db}
}

func (repo userRepository) Insert(user entity.User) (int64, error) {
	var id int64
	err := repo.db.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo userRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	err := repo.db.QueryRow("SELECT id, username, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (repo userRepository) GetById(id string) (entity.User, error) {
	var user entity.User
	err := repo.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

