package main

import (
	"crypto/rand"
	"encoding/hex"

	lhttp "github.com/alireza-aliabadi/golang-hexagonal/internal/library/adapters/http"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/library/adapters/repo"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/library/core"

	"github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type randID struct{}
func randHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func (randID) BookID() core.BookID {
	return core.BookID("book-", randHex(8))
}
func (randID) UserID() core.UserID {
	return core.UserID("user-", randHex(8))
}

func main() {
	books := repo.NewMemBooks()
	users := repo.NewMemUsers()
	svc := core.NewBookService(books, users, randID{})

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger(), middleware.Recover())

	h := lhttp.NewHandler(svc)
	h.Register(e)

	e.Logger.Fatal(e.Start(":8081"))
}