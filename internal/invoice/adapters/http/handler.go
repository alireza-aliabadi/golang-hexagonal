package http

import (
	"net/http"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/invoice/core"
	"github.com/labstack/echo/v4"
)

type Handler struct{
	svc *core.InvoiceService
}

// builder function
func NewHandler(s *core.InvoiceService) *Handler {
	return &Handler{
		svc: s,
	}
}

// echo handlers register
func (h *Handler) Register(e *echo.Echo) {
	e.POST("/invoices", h.create)
	e.GET("/invoices", h.list)
	e.POST("/invoices/:id/pay", h.pay)
}

// implement handlers
func (h *Handler) create(c echo.Context) error {
	var req struct{
		Client string `json:"client"`
		Amount int64 `json:"amount"`
	}
	if err := c.Bind(&req); err != nil || req.Client == "" || req.Amount <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "client ans positive amount are required.")
	}
	invoice, err := h.svc.Create(c.Request().Context(), req.Client, req.Amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, invoice)
}
func (h *Handler) list(c echo.Context) error {
	onlyUnpaid := c.QueryParam("unpaid") == "true"
	invoice, err := h.svc.List(c.Request().Context(), onlyUnpaid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch data.")
	}
	return c.JSON(http.StatusOK, invoice)
}
func (h *Handler) pay(c echo.Context) error {
	id := core.InvoiceID(c.Param("id"))
	invoice, err := h.svc.MarkPaid(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, invoice)
}