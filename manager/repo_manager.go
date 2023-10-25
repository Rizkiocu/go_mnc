package manager

import "test_mnc/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	PaymentRepo() repository.PaymentRepository
}

type repoManager struct {
	infraManager InfraManager
}

// PaymentRepo implements RepoManager.
func (r *repoManager) PaymentRepo() repository.PaymentRepository {
	return repository.NewPaymentRepository(r.infraManager.Conn())
}

// UseRepo implements RepoManager.
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infraManager.Conn())
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
