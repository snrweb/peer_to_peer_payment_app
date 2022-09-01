package models

import (
	"errors"

	"github.com/google/uuid"
)

type Account struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id"`
	CurrencyType int     `json:"currency_type"`
	Balance      float64 `json:"balance"`
}

var (
	accounts = map[string]*Account{"b7392d0b-8a6f-4436-869a-037d054ea7d5": {
		ID:           "d78002d0b-9a6f-4436-999a-237d054ea7d5",
		UserID:       "b7392d0b-8a6f-4436-869a-037d054ea7d5",
		CurrencyType: 1,
		Balance:      99.0,
	}, "d78002d0b-9a6f-4436-999a-237d054ea7d5": {
		ID:           "b7392d0b-8a6f-4436-869a-037d054ea7d5",
		UserID:       "d78002d0b-9a6f-4436-999a-237d054ea7d5",
		CurrencyType: 1,
		Balance:      6.0,
	}}
)

func Create(userAcounts []Account) error {
	for _, account := range userAcounts {
		account.ID = uuid.New().String()

		_, isAvailable := accounts[account.UserID]
		if isAvailable {
			return errors.New("Account is available")
		}

		accounts[account.UserID] = &account
		accounts[account.UserID].Balance = 0.0
	}

	return nil
}

func (account *Account) GetBalance() (float64, error) {
	_, isAvailable := accounts[account.UserID]
	if !isAvailable {
		return 0.0, errors.New("Account not available")
	}

	return accounts[account.UserID].Balance, nil
}

func (account *Account) Deposit() (Account, error) {
	_, isAvailable := accounts[account.UserID]
	if !isAvailable {
		return Account{}, errors.New("Account not available")
	}

	accounts[account.UserID].Balance += account.Balance
	return *accounts[account.UserID], nil
}

func (account *Account) Transfer(user User) (Account, error) {
	_, isAvailable := accounts[account.UserID]
	if !isAvailable {
		return Account{}, errors.New("Account not available")
	}

	_, isAvailable = accounts[user.ID]
	if !isAvailable {
		return Account{}, errors.New("Recipient account not available")
	}

	if accounts[account.UserID].Balance < account.Balance {
		return Account{}, errors.New("Insufficient funds")
	}

	accounts[account.UserID].Balance -= account.Balance
	accounts[user.ID].Balance += account.Balance
	return *accounts[account.UserID], nil
}

func (account *Account) Withdraw() (Account, error) {
	_, isAvailable := accounts[account.UserID]
	if !isAvailable {
		return Account{}, errors.New("Account not available")
	}

	if accounts[account.UserID].Balance < account.Balance {
		return Account{}, errors.New("Insufficient funds")
	}

	accounts[account.UserID].Balance -= account.Balance
	return *accounts[account.UserID], nil
}
