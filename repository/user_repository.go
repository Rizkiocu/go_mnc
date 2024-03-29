package repository

import (
	"database/sql"
	"test_mnc/model"
)

type UserRepository interface {
	Save(payload model.UserCredential) error
	FindByEmail(email string) (model.UserCredential, error)
	FindById(id string) (model.UserCredential, error)
}

type userRepository struct {
	db *sql.DB
}

// FindByEmailimplements UserRepository.
func (u *userRepository) FindByEmail(email string) (model.UserCredential, error) {
	row := u.db.QueryRow("SELECT id, email, password, name FROM user_credential WHERE email = $1 AND is_active = $2", email, true)
	var userCredential model.UserCredential
	err := row.Scan(&userCredential.Id, &userCredential.Email, &userCredential.Password, &userCredential.Name)
	if err != nil {
		return model.UserCredential{}, err
	}
	return userCredential, nil
}

// Save implements UserRepository.
func (u *userRepository) Save(payload model.UserCredential) error {
	_, err := u.db.Exec("INSERT INTO user_credential VALUES ($1, $2, $3, $4, $5)", payload.Id, payload.Email, payload.Password, payload.Name, true)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindById(id string) (model.UserCredential, error) {
	row := u.db.QueryRow("SELECT id, email, password, name FROM user_credential WHERE id = $1", id)
	var user model.UserCredential
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name)
	if err != nil {
		return model.UserCredential{}, err
	}
	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
