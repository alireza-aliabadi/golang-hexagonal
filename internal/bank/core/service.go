package core

import (
	"context"
	"errors"
	"fmt"
	"strings"
)


// define custom errors
var (
	ErrInvalidName = errors.New("owner name shouldn't be empty")
	ErrInvalidAmount = errors.New("amount must be positive")
	ErrTaNotFound = errors.New("could't found any Transaction")
	ErrInsufficient = errors.New("insufficient funds")
)
func ErrAccNotFound(AccountID AccountID) error {
	return fmt.Errorf("couldn't find any Account by id: %s", AccountID)
}

// service to deploy banking functionality
type BankingService struct {
	accounts AccountReport
	tas TaReport
	clock Clock
	ids GenID
}

// service builder
func NewBankingService(a AccountReport, t TaReport, clock Clock, ids GenID) *BankingService {
	return &BankingService{
		accounts: a,
		tas: t,
		clock: clock,
		ids: ids
	}
}

func (s *BankingService) CreateAccount(ctx context.Context, owner string) (*Account, error) {
	owner = strings.TrimSpace(owner)
	if owner == "" {
		return nil, ErrInvalidName
	}
	acc := &Account {
		ID: s.ids.AccountID(),
		Owner: owner,
		Balance: 0
	}
	if err := s.accounts.Save(ctx, acc); err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *BankingService) Deposit(ctx context.Context, to AccountID, amount int64, note string) (*Transaction, *Account, error) {
	if amount <= 0 {
		return nil, nil, ErrInvalidAmount
	}
	toAcc, err := s.accounts.ByID(ctx, to)
	if err != nil {
		return nil, nil, ErrAccNotFound(to)
	}
	toAcc.Recieve(amount)
	if err:= s.accounts.Save(ctx, toAcc); err != nil {
		return nil, nil, err
	}
	ta := &Transaction {
		ID: s.ids.TaID(),
		To: &to,
		Amount: amount,
		Note: strings.TrimSpace(note),
		CreatedAt: s.clock.Now()
	}
	if err := s.tas.Add(ctx, ta); err != nil {
		return nil, nil, err
	}
	return ta, toAcc, nil
}

func (s *BankingService) Transfer(ctx context.Context, from, to AccountID, amount int64, note string) (*Transaction, *Account, *Account, error) {
	if amount <= 0 {
		return nil, nil, nil, ErrInvalidAmount
	}
	src, err := s.accounts.ByID(ctx, from)
	if err != nil {
		return nil, nil, nil, ErrAccNotFound(from)
	}
	dst, err := s.accounts.ByID(ctx, to)
	if err != nil {
		return nil, nil, nil, ErrAccNotFound(to)
	}
	if src.Balance < amount {
		return nil, nil, nil, ErrInsufficient
	}
	src.Pay(amount)
	dst.Recieve(amount)
	if err := s.accounts.Save(ctx, src); err != nil {
		return nil, nil, nil, err
	}
	if err := s.accounts.Save(ctx, dst); err != nil {
		return nil, nil, nil, err
	}
	ta := &Transaction {
		ID: s.ids.TaID(),
		From: &from,
		To: &to,
		Amount: amount
		Note: strings.TrimSpace(note),
		CreatedAt: s.clock.Now()
	}
	if err := s.tas.Add(ctx, ta); err != nil {
		return nil, nil, nil, err
	}
	return ta, src, dst, nil
}

func (s *BankingService) Transactions(ctx context.Context, accID AccountID) ([]*Transaction, error) {
	return s.tas.ByAccount(ctx, accID)
}

func (s *BankingService) Accounts(ctx context.Context) ([]*Account, err) {
	return s.accounts.List(ctx)
}
