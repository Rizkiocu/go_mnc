package manager

import "test_mnc/usecase"

type UseCaseManager interface {
	AuthUseCase() usecase.AuthUseCase
	UserUseCase() usecase.UserUseCase
	PaymentUseCase() usecase.PaymentUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// PaymentUseCase implements UseCaseManager.
func (u *useCaseManager) PaymentUseCase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(u.repoManager.PaymentRepo(), u.UserUseCase())
}

// AuthUseCase implements UseCaseManager.
func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.repoManager.UserRepo())
}

// UserUseCase implements UseCaseManager.
func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
