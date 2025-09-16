package http

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/alireza-aliabadi/golang-hexagonal/internal/library/core
"
)

type Handler struct{
	svc *core.BookService
}

// builder function
func NewHandler(s *core.BookService) *Handler {
	return &Handler{
		svc: s
	}
}

// echo routes registeration method
func (h *Handler) Register(e *echo.Echo) {
	e.Post("/books", h.addBook)
	e.GET("/books", h.listBooks)
	e.POST("/users", h.addUser)
	e.GET("/users", h.listUsers)
	e.POST("/borrow", h.borrow)
	e.POST("/return/:bookID", h.returnBook)
}

// implemet handler functions
func (h *Handler) addBook(c echo.Context) error {
	var req struct{
		Title string
		Author string
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	b, err := h.svc.AddBook(c.Request().Context(), req.Title, req.Author)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, b)
}

func (h *Handler) listBooks(c echo.Context) error {
	books, err := h.svc.ListBooks(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch data")
	}
	return c.JSON(http.StatusOK, books)
}

func (h * Handler) addUser(c echo.Context) error {
	var req struct{
		Name string
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	u, err := h.svc.RegisterUser(c.Request().Context(), req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *Hnadler) listUsers(c echo.Context) error {
	users, err := h.svc.ListUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch data")
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) borrrow(c echo.Context) error {
	var req struct{
		UserID: string,
		BookID: string
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	b, err := h.svc.BorrowBook(c.Request().Context(), core.UserID(req.UserID), core.BookID(req.BookID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, b)
}

func (h *Handler) returnBook(c echo.Context) error {
	BookID := core.BookID(c.Param("bookID"))
	b, err := h.svc.ReturnBook(c.Request().Context(), BookID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, b)
}