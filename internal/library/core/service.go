package core

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrInvalid = errors.New("invalid input")
	ErrNotFound = errors.New("not found")
	ErrUnavailable = errors.New("book not available")
)

type BookService struct {
	books BookRepo
	users UserRepo
	ids IDGen
}

// builder function
func NewBookService(br BookRepo, ur UserRepo, ids IDGen) *BookService {
	return &BookService{
		books: br,
		users: ur,
		ids: ids,
	}
}

func (bs *BookService) AddBook(ctx context.Context, title, author string) (*Book, error) {
	title, author = strings.TrimSpace(title), strings.TrimSpace(author)
	if title == "" || author == "" {
		return nil, ErrInvalid
	}
	book := &Book{
		ID: bs.ids.BookID(),
		Title: title,
		Author: author,
		Status: "availabel",
	}
	if err := bs.books.Save(ctx, book); err != nil {
		return nil, err
	}
	return book, nil
}

func (bs *BookService) RegisterUser(ctx context.Context, name string) (*User, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrInvalid
	}
	user := &User{
		ID: bs.ids.UserID(),
		Name: name,
	}
	if err := bs.users.Save(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (bs *BookService) BorrowBook(ctx context.Context, userID UserID, bookID BookID) (*Book, error) {
	user, err := bs.users.ByID(ctx, userID)
	if err != nil || user == nil {
		return nil, ErrNotFound
	}
	book, err := bs.books.ByID(ctx, bookID)
	if err != nil || book == nil {
		return nil, ErrNotFound
	}
	if book.Status != "available" {
		return nil, ErrUnavailable
	}
	book.Status = "borrowed"
	book.Borrower = &user.ID
	if err := bs.books.Save(ctx, book); err != nil {
		return nil, err
	}
	return book, nil
}

func (bs *BookService) ReturnBook(ctx context.Context, bookID BookID) (*Book, error) {
	book, err := bs.books.ByID(ctx, bookID)
	if err != nil {
		return nil, ErrNotFound
	}
	book.Status = "available"
	book.Borrower = nil
	if err := bs.books.Save(ctx, book); err != nil {
		return nil, err
	}
	return book, nil
}

func (bs *BookService) ListBooks(ctx context.Context) ([]*Book, error) {
	return bs.books.List(ctx)
}

func (bs *BookService) ListUsers(ctx context.Context) ([]*User, error) {
	return bs.users.List(ctx)
}