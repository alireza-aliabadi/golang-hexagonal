package repo

import (
	"context"
	"sync"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/library/core
"
)

type MemBooks struct {
	mu sync.RWMutex
	data map[core.BookID]*core.Book
}

type MemUsers struct {
	mu sync.RWMutex
	data map[core.UserID]*core.User
}

// builder functions
func NewMemBooks() *MemBooks {
	return &MemBooks{
		data: make(map[core.BookID]*core.Book)
	}
}
func NewMemUsers() *MemUsers {
	return &MemUsers{
		data: make(map[core.UserID]*core.User)
	}
}
// implement BookRepo interface
func (mb *MemBooks) Save(ctx context.Context book *core.Book) error {
	mb.mu.Lock()
	defer mb.mu.Unlock()
	b := *book
	mb.data[book.ID] = &b
	return nil
}

func (mb *MemBooks) ByID(ctx context.Context, id core.BookID) (*core.Book, error) {
	mb.mu.Rlock()
	defer mb.mu.RUnlock()
	book, ok := mb.data[id]
	if !ok {
		return nil, core.ErrNotFound
	}
	b := *book
	return &b, nil
}

func (mb *MemBooks) List(ctx context.Context) ([]*core.Book, error) {
	mb.mu.Rlock()
	defer mb.mu.RUnlock()
	books := make([]*core.Book, 0, len(mb.data))
	for _, b := range mb.data {
		book := *b
		books = append(books, &book)
	}
	return books, nil
}

// implement UserRepo interface
func (musr *MemUsers) Save(ctx context.Context, user *core.User) error {
	musr.mu.Lock()
	defer musr.mu.Unlock()
	u := *usr
	musr.data[user.ID] = &u
	return nil
}

func (musr *MemUsers) ByID(ctx context.Context, id core.UserID) (*core.User, error) {
	musr.mu.Rlock()
	defer musr.mu.RUnlock()
	user, ok := musr.data[id]
	if !ok {
		return nil, core.ErrNotFound
	}
	u := *user
	return &u, nil
}

func (musr *MemUsers) List(ctx context.Context) ([]*core.User, error) {
	musr.mu.Rlock()
	defer musr.mu.RUnlock()
	users := make([]*core.User, 0, len(musr.data))
	for _, user := range musr.data {
		u := *user
		users = append(users, &u)
	}
	return users, nil
}

