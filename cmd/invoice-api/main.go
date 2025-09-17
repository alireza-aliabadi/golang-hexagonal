package main

import (
	"crypto/rand"
	"encoding/hex"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	invhttp "github.com/alireza-aliabadi/golang-hexagonal/internal/invoice/adapters/http"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/invoice/adapters/repo"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/invoice/core"	
)

type sysClock struct{}
func (sysClock) Now() time.Time {
	return time.Now().UTC()
}

func randHex(n int) string { 
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
type randID struct{}
func (randID) InvoiceID() core.InvoiceID {
	return core.InvoiceID("invoice-" + randHex(8))
}

func main() {
	r := repo.NewMemRepo()
	svc := core.NewInvoiceService(r, sysClock{}, randID{})

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger(), middleware.Recover())

	h := invhttp.NewHandler(svc)
	h.Register(e)

	e.Logger.Fatal(e.Start(":8082"))
}