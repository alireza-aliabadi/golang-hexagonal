package core

import "time"

type AccountID string
// TaID -> Transaction ID
type TaID string

type Account struct {
	ID AccountID `json:"id"`
	Owner string `json:"owner"`
	Balance int64 `json:"balance"`
}

type Transaction struct {
	ID TaID `json:"id"`
	From *AccountID `json:"from,omitempty"`
	To *AccountID `json:"to,omitempty"`
	Amount int64 `json:"amount"`
	Note string `json:"note,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Account) Recieve(amount int64) {
	a.Balance += amount
}

func (a *Account) Pay(amount int64) {
	a.Balance -= amount
} 