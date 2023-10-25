package usecase

import (
	"errors"
	"fmt"
	"regexp"
	"test_mnc/model"
	"test_mnc/model/dto"
	"test_mnc/repository"
	"test_mnc/util/common"
	"test_mnc/util/security"
)

type UserUseCase interface {
	FindByEmail(email string) (model.UserCredential, error)
	Register(payload dto.AuthRequest) error
	FindById(id string) (model.UserCredential, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// FindByUsername implements UserUseCase.
func (u *userUseCase) FindByEmail(email string) (model.UserCredential, error) {
	return u.repo.FindByEmail(email)
}

func (u *userUseCase) FindById(id string) (model.UserCredential, error) {
	user, err := u.repo.FindById(id)
	if err != nil {
		return model.UserCredential{}, fmt.Errorf("uom not found")
	}
	return user, nil
}

// Register implements UserUseCase.
func (u *userUseCase) Register(payload dto.AuthRequest) error {
	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}
	// Check if email is valid (e.g., gmail.com)
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	validEmail := regexp.MustCompile(emailPattern)
	if !validEmail.MatchString(payload.Email) {
		return errors.New("invalid email")
	}

	//password requirement area
	if len(payload.Password) < 6 {
		return fmt.Errorf("password must contain at least six number")
	}
	userCredential := model.UserCredential{
		Id:       common.GenerateUUID(),
		Email:    payload.Email,
		Password: hashPassword,
		Name:     payload.Name,
	}

	err = u.repo.Save(userCredential)
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}
	return nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}
