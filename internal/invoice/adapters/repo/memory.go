package repo

import (
	"context"
	"sync"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/invoice/core"
)

type MemRepo struct{
	mu sync.RWMutex
	data map[core.InvoiceID]*core.Invoice
}

// builder function
func NewMemRepo() *MemRepo {
	return &MemRepo{
		data: make(map[core.InvoiceID]*core.Invoice)
	}
}

// implement Repo interface
func (mr *MemRepo) Save(ctx context.Context, invc *core.Invoice) error {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	invoice := *invc
	mr.data[invc.ID] = &invoice
	return nil
}

func (mr *MemRepo) ByID(ctx context.Context, id core.InvoiceID) (*core.Invoice, error) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	invoice, ok := mr.data[id]
	if !ok {
		return nil, core.ErrNotFound
	}
	invc := *invoice
	return &invc, nil
}

func (mr *MemRepo) List(ctx context.Context, onlyUnpaid bool) ([]*core.Invoice, error) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	result := []*core.Invoice{}
	for _, invoice := range mr.data {
		if onlyUnpaid && invoice.Paid {
			invc := *invoice
			result = append(result, &invc)
		}
	}
	return result, nil
}

