package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, card CreditCard) error
	GetCreditCard(card CreditCard) (CreditCard, error)
	CreateCreditCard(card CreditCard) error
}

type Transaction struct {
	ID          string
	Amount      float64
	Status      string
	Description string
	Store       string
	CreditCard  string
	CreatedAt   time.Time
}

func NewTransaction() *Transaction {
	transaction := &Transaction{}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	return transaction
}

func (t *Transaction) ProccessAndValidate(card *CreditCard) {
	if t.Amount+card.Balance > card.Limit {
		t.Status = "Rejected"
	} else {
		t.Status = "Approved"
		card.Balance = card.Balance + t.Amount
	}
}
