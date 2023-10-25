package model

import "time"

type Payment struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Purchase_date  time.Time      `json:"purchase_date"`
	Price          int            `json:"price"`
	UserCredential UserCredential `json:"user_credential"`
}
