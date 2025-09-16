package core

import "time"

type InvoiceID string

type Invoice struct{
	ID InvoiceID `json:"id"`
	Client string `json:"client"`
	Amount int64 `json:"amount"`
	Paid bool `json:"paid"`
	CreatedAt time.Time `json:"created_at"`
	PaidAt *time.Time `json:"paid_at,omitempty"`
}

func (iv *Invoice) MarkPaid(t time.Time) {
	if !iv.Paid {
		iv.Paid = true
		iv.PaidAt = &t
	}
}

