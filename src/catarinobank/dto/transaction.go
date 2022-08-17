package dto

import "time"

type Transaction struct {
	ID              string
	CreatedAt       time.Time
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Amount          float64
	Store           string
	Description     string
}
