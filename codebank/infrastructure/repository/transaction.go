package repository

import (
	"database/sql"

	"github.com/Valdeci97/fullcycle-event/domain"
)

type TransctionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransctionRepositoryDb {
	return &TransctionRepositoryDb{}
}

func (t *TransctionRepositoryDb) Save(
	transaction *domain.Transaction,
	card domain.CreditCard,
) error {
	stmt, err := t.db.Prepare(`
		insert into transactions (id, credit_card, amount, status, description, store, created_at)
		values ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCard,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	)
	if err != nil {
		return err
	}
	if transaction.Status == "Approved" {
		err = t.UpdateBalance(card)
		if err != nil {
			return err
		}
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (t *TransctionRepositoryDb) UpdateBalance(card domain.CreditCard) error {
	_, err := t.db.Exec("update balances set balance = $1 where id = $2", card.Balance, card.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransctionRepositoryDb) CreateCreditCard(card domain.CreditCard) error {
	stmt, err := t.db.Prepare(`
		inser into credit_cards (
			id,
			name,
			number,
			expiration_month,
			expiration_year,
			cvv,
			balance,
			balance_limit
		) values ($1, $2, $3, $4, $5, $6, $7, $8
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		card.ID,
		card.Name,
		card.Number,
		card.ExpirationMonth,
		card.ExpirationYear,
		card.CVV,
		card.Balance,
		card.Limit,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return nil
	}
	return nil
}
