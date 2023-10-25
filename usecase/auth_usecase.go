package usecase

import (
	"fmt"
	"test_mnc/model/dto"
	"test_mnc/repository"
	"test_mnc/util/security"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequest) (dto.AuthResponse, error)
	Logout(email string) error
}

type authUseCase struct {
	repo repository.UserRepository
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(payload dto.AuthRequest) (dto.AuthResponse, error) {
	// Username di db
	user, err := a.repo.FindByEmail(payload.Email)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("unauthorized: invalid credential")
	}
	// Validasi Password
	err = security.VerifyPassword(user.Password, payload.Password)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("unauthorized: invalid credential")
	}

	// Generate Token
	token, err := security.GenerateJwtToken(user)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	return dto.AuthResponse{
		Email: user.Email,
		Token: token,
	}, nil
}

func (a *authUseCase) Logout(email string) error {

	return nil
}

func NewAuthUseCase(repo repository.UserRepository) AuthUseCase {
	return &authUseCase{repo: repo}
}
