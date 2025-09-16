package core

import "context"

type BookRepo interface {
	Save(ctx context.Context, b *Book) error
	ByID(ctx context.Context, id BookID) (*Book, error)
	List(ctx context.Context) ([]*Book, error)
}

type UserRepo interface {
	Save(ctx context.Context, u *User) error
	ByID(ctx context.Context, id UserID) (*User, error)
	List(ctx context.Context) ([]*User, error)
}

type IDGen interface {
	BookID() BookID
	UserID() UserID
}