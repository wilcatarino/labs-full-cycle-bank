package repository

import (
	"database/sql"
	"errors"

	"github.com/wilcatarino/labs-full-cycle-bank/domain"
)

type TransactionRepositoryDatabase struct {
	db *sql.DB
}

func NewTransactionRepositoryDatabase(db *sql.DB) *TransactionRepositoryDatabase {
	return &TransactionRepositoryDatabase{db: db}
}

func (transactionRepositoryDatabase *TransactionRepositoryDatabase) updateBalance(creditCard domain.CreditCard) error {
	_, err := transactionRepositoryDatabase.db.Exec("update credit_cards set balance = $1 where id = $2", creditCard.Balance, creditCard.ID)
	if err != nil {
		return err
	}
	return nil
}

func (transactionRepositoryDatabase *TransactionRepositoryDatabase) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := transactionRepositoryDatabase.db.Prepare(`
		insert into transactions(id, created_at, credit_card_id, amount, status, description, store)
		values($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreatedAt,
		transaction.CreditCardId,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
	)
	if err != nil {
		return err
	}
	if transaction.Status == "approved" {
		err = transactionRepositoryDatabase.updateBalance(creditCard)
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

func (transactionRepositoryDatabase *TransactionRepositoryDatabase) CreateCreditCard(creditCard domain.CreditCard) error {
	stmt, err := transactionRepositoryDatabase.db.Prepare(`
		insert into credit_cards(id, created_at, name, number, expiration_month, expiration_year, CVV, balance, balance_limit) 
		values($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.CreatedAt,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDatabase) GetCreditCard(creditCard domain.CreditCard) (domain.CreditCard, error) {
	var c domain.CreditCard
	stmt, err := t.db.Prepare("select id, balance, balance_limit from credit_cards where number = $1")
	if err != nil {
		return c, err
	}
	if err = stmt.QueryRow(creditCard.Number).Scan(&c.ID, &c.Balance, &c.Limit); err != nil {
		return c, errors.New("credit card does not exists")
	}
	return c, nil
}
