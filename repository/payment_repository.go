package repository

import (
	"database/sql"
	"test_mnc/model"
)

type PaymentRepository interface {
	Save(payment model.Payment) error
	FindById(id string) (model.Payment, error)
	FindAll() ([]model.Payment, error)
	Update(payment model.Payment) error
	DeleteById(id string) error
}

type paymentRepository struct {
	db *sql.DB
}

// DeleteById implements PaymentRepository.
func (p *paymentRepository) DeleteById(id string) error {
	_, err := p.db.Exec("DELETE FROM payment WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements PaymenyRepository.
func (p *paymentRepository) FindAll() ([]model.Payment, error) {
	rows, err := p.db.Query(`SELECT p.id, p.name, p.purchase_date ,p.price, u.id, u.name FROM payment p
	JOIN user_credential u ON u.id = p.user_credential_id`)
	if err != nil {
		return nil, err
	}
	var payments []model.Payment
	for rows.Next() {
		payment := model.Payment{}
		err := rows.Scan(
			&payment.Id,
			&payment.Name,
			&payment.Purchase_date,
			&payment.Price,
			&payment.UserCredential.Id,
			&payment.UserCredential.Name,
		)
		if err != nil {
			return nil, err
		}
	}
	return payments, nil
}

// FindById implements PaymenyRepository.
func (p *paymentRepository) FindById(id string) (model.Payment, error) {
	row := p.db.QueryRow(`SELECT p.id, p.name,p.purchase.date, p.price, u.id, u.name FROM payment p
	JOIN user_credential u ON u.id = p.user_credential_id WHERE p.id = $1`, id)
	payment := model.Payment{}
	err := row.Scan(
		&payment.Id,
		&payment.Name,
		&payment.Purchase_date,
		&payment.Price,
		&payment.UserCredential.Id,
		&payment.UserCredential.Name,
	)
	if err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}

// Save implements PaymentRepository.
func (p *paymentRepository) Save(payment model.Payment) error {
	_, err := p.db.Exec("INSERT INTO payment VALUES ($1, $2, $3, $4, $5)",
		payment.Id,
		payment.Name,
		payment.Purchase_date,
		payment.Price,
		payment.UserCredential.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

// Update implements PaymentRepository.
func (p *paymentRepository) Update(payment model.Payment) error {
	_, err := p.db.Exec("UPDATE payment SET name = $2, purchase_date = $3, price = $4, uom_id = $5 WHERE id = $1",
		payment.Id,
		payment.Name,
		payment.Purchase_date,
		payment.Price,
		payment.UserCredential.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db: db}
}
