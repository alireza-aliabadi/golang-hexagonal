package core

import (
	"context"
	"strings"
	"errors"
)

// custom errors
var (
	ErrInvalid = errors.New("invalid input")
	ErrNotFound = errors.New("invoice not found")
)

type InvoiceService struct{
	repo Repo
	clck Clock
	ids IDGen
}

// builder function
func NewInvoiceService(r Repo, c Clock, ids IDGen) *InvoiceService {
	return &InvoiceService{
		repo: r,
		clck: c,
		ids: ids,
	}	
}

// methods implementation
func (svc *InvoiceService) Create(ctx context.Context, client string, amount int64) (*Invoice, error) {
	client = strings.TrimSpace(client)
	if client == "" || amount <= 0 {
		return nil, ErrInvalid
	}
	invoice := &Invoice{
		ID: svc.ids.InvoiceID(),
		Client: client,
		Amount: amount,
		Paid: false,
		CreatedAt: svc.clck.Now(),
	}
	if err := svc.repo.Save(ctx, invoice); err != nil {
		return nil, err
	}
	return invoice, nil
}

func (svc *InvoiceService) MarkPaid(ctx context.Context, id InvoiceID) (*Invoice, error) {
	invoice, err := svc.repo.ByID(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}
	invoice.MarkPaid(svc.clck.Now())
	if err := svc.repo.Save(ctx, invoice); err != nil {
		return nil, err
	}
	return invoice, nil
}

func (svc *InvoiceService) List(ctx context.Context, onlyUnpaid bool) ([]*Invoice, error) {
	return svc.repo.List(ctx, onlyUnpaid)
}