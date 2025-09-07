package core

import (
	"time"
	"context"
)

type AccountReport interface {
	Save(ctx context.Context, a *Account) error
	List(ctx context.Context) ([]*Account, error)
	ByID(ctx context.Context, id AccountID) (a *Account, error)  
}

type TaReport interface {
	ByAccount(ctx context.Context, id AccountID) ([]*Transaction, error)
	Add(ctx context.Context t *Transaction) error
}

type Clock interface {
	Now() time.Time
}

type GenID interface {
	AccountID() AccountID
	TaID() TaID
}