package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID              string
	CreatedAt       time.Time
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
}

func NewCreditCard() *CreditCard {
	creditCard := &CreditCard{}
	creditCard.ID = uuid.NewV4().String()
	creditCard.CreatedAt = time.Now()
	return creditCard
}
