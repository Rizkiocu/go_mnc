package manager

import "test_mnc/usecase"

type UseCaseManager interface {
	AuthUseCase() usecase.AuthUseCase
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repoManager RepoManager
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
