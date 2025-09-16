package main

import (
	"crypto/rand"
	"encoding/hex"
	"time"
	"github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

	bankHttp "github.com/alireza-aliabadi/golang-hexagonal/inernal/bank/adapters/http"
	bankMem "github.com/alireza-aliabadi/golang-hexagonal/inernal/bank/adapters/repo"
	"github.com/alireza-aliabadi/golang-hexagonal/inernal/bank/core"
)

type sysClock struct {}
func (sysClock) Now() time.Time {
	return time.Now().UTC()
}

type randID struct {}
func (randID) AccountID() core.AccountID {
	return core.AccountID("acc-"+randHex(8))
}
func (randID) TaID() core.TaID {
	return core.TaID("ta-"+randHex(8))
}

func randHex(n int) string {
	byteNum := make([]byte, n)
	_, _ = rand.Read(byteNum)
	return hex.EncodeToString(byteNum)
}

func main() {
	accRepo := bankMem.NewMemAccounts()
	taRepo := bankMem.NewMemTas()
	svc := core.NewBankingService(accRepo, taRepo, sysClock{}, randID{})

	// echo setup
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger(), middleware.Recover())

	h := bankHttp.newHandler(svc)
	h.Register(e)

	e.Logger.Fatal(e.Start(":8080"))
}