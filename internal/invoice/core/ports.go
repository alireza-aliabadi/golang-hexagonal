package core

import (
	"context"
	"time"
)

type Clock interface{
	Now() time.Time
}

type Repo interface{
	Save(ctx context.Context, iv *Invoice) error
	ByID(ctx context.Context, id InvoiceID) (*Invoice, error)
	List(ctx context.Context, onlyUnpaid bool) ([]*Invoice, error)
}

type IDGen interface{
	InvoiceID() InvoiceID
}