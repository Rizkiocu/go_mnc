package usecase

import (
	"fmt"
	"test_mnc/model"
	"test_mnc/repository"
	"test_mnc/util/common"
)

type PaymentUseCase interface {
	CreateNew(payload model.Payment) error
	FindById(id string) (model.Payment, error)
	FindAll() ([]model.Payment, error)
	Update(payload model.Payment) error
	Delete(id string) error
}

type paymentUseCase struct {
	repo   repository.PaymentRepository
	userUC UserUseCase
}

// CreateNew implements PaymentUseCase.
func (p *paymentUseCase) CreateNew(payload model.Payment) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.Price <= 0 {
		return fmt.Errorf("price must be greater than zero")
	}

	payload.Id = common.GenerateUUID()
	_, err := p.userUC.FindById(payload.UserCredential.Id)
	if err != nil {
		return err
	}
	err = p.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to save new Payment: %v", err)
	}

	return nil
}

// Delete implements PaymentUseCase.
func (p *paymentUseCase) Delete(id string) error {
	payment, err := p.FindById(id)
	if err != nil {
		return err
	}

	err = p.repo.DeleteById(payment.Id)
	if err != nil {
		return fmt.Errorf("failed to delete Payment: %v", err)
	}

	return nil
}

// FindAll implements PaymentUseCase.
func (p *paymentUseCase) FindAll() ([]model.Payment, error) {
	return p.repo.FindAll()
}

// FindById implements PaymentUseCase.
func (p *paymentUseCase) FindById(id string) (model.Payment, error) {
	payment, err := p.repo.FindById(id)
	if err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}

// Update implements PaymentUseCase.
func (p *paymentUseCase) Update(payload model.Payment) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.Price <= 0 {
		return fmt.Errorf("price must be greater than zero")
	}

	_, err := p.userUC.FindById(payload.UserCredential.Id)
	if err != nil {
		return err
	}

	_, err = p.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = p.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update Payment: %v", err)
	}

	return nil
}

func NewPaymentUseCase(repo repository.PaymentRepository, userUc UserUseCase) PaymentUseCase {
	return &paymentUseCase{
		repo:   repo,
		userUC: userUc,
	}
}
