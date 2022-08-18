package usecase

import (
	"time"

	"github.com/Valdeci97/fullcycle-event/domain"
	"github.com/Valdeci97/fullcycle-event/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func NewUseCaseTransaction(repository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{repository}
}

func (use UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	card := use.HydrateCard(transactionDto)
	cardBalanceAndLimit, err := use.TransactionRepository.GetCreditCard(*card)
	if err != nil {
		return domain.Transaction{}, err
	}
	card.ID = cardBalanceAndLimit.ID
	card.Limit = cardBalanceAndLimit.Limit
	card.Balance = cardBalanceAndLimit.Balance
	transaction := use.NewTransaction(transactionDto, cardBalanceAndLimit)
	transaction.ProccessAndValidate(card)
	err = use.TransactionRepository.Save(*transaction, *card)
	if err != nil {
		return domain.Transaction{}, err
	}
	return *transaction, nil
}

func (use UseCaseTransaction) HydrateCard(dto dto.Transaction) *domain.CreditCard {
	card := domain.NewCreditCard()
	card.Name = dto.Name
	card.Number = dto.Number
	card.ExpirationMonth = dto.ExpirationMonth
	card.ExpirationYear = dto.ExpirationYear
	card.CVV = dto.CVV
	return card
}

func (use UseCaseTransaction) NewTransaction(transaction dto.Transaction, card domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.ID = transaction.ID
	t.Amount = transaction.Amount
	t.Store = transaction.Store
	t.Description = transaction.Description
	t.CreatedAt = time.Now()
	return t
}
