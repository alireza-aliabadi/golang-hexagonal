package http

import (
	"fmt"
	"net/http"
	"time"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/bank/core"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *core.BankingService
}

// creator function
func newHandler(s *core.BankingService) *Handler {
	return &Handler{svc: s}
}

//register routes and their handlers
func (h *Handler) Register(e *echo.Echo) {
	e.POST("/accounts", h.createAccount)
	e.GET("/accounts", h.listAccounts)
	e.POST("/accounts/:id/deposit", h.deposit)
	e.POST("/transfer", h.transfer)
	e.GET("/accounts/:id/transactions", h.transactions)
}

func (h *Handler) createAccount(c echo.Context) error {
	var body struct {
		Owner string `json:"owner"`
	}
	if err := c.Bind(&body); err != nil || body.Owner == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "owner isn't provided")
	}
	acc, err := h.svc.CreateAccount(c.Request().Context(), body.owner)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, acc)
}


func (h *Handler) listAccounts(c echo.Context) error {
	accs, err := h.svc.Accounts(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed")
	}
	return c.JSON(http.StatusOK, accounts)
}

func (h *Handler) deposit(c echo.Context) error {
	var body struct {
		Amount int64 `json:"amount"`
		Note string `json:"note"`
	}

	id := core.AccountID(c.Param("id"))
	if err := c.Bind(&body); err != nil || body.Amount <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "positive amount required")
	}
	ta, acc, err := h.svc.Deposit(c.Request().Context(), id, body.Amount, body.Note)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"transaction": ta,
		"account": acc,honnamkuan.golang-snippets 
		"when": time.Now().UTC(),
	})
}

func (h *Handler) transfer(c echo.Context) error {
	var body struct {
		From string `json:"from"`
		To string `json:"to"`
		Amount int64 `json:"amount"`
		Note string `json:"note"`
	}
	if err := c.Bind(&body); err != nil || body.Amount <= 0 || body.From == "" || body.To == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "positive amount, from and to are required")
	}
	ta, src, dst, err := h.svc.Transfer(c.Request().Context(), core.AccountID(body.From), core.AccountID(body.To), body.Amount, body.Note)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"transaction": ta,
		"from": src,
		"to": dst,
	})
}

func (h *Handler) transactions(c echo.Context) error {
	id := core.AccountID(c.Param("id"))
	items, err := h.svc.Transactions(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to fetch transactions by: %v", err))
	}
	return c.JSON(http.StatusOK, items)
}