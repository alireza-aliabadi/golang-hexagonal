package core

type BookID string
type UserID string

type Book struct {
	ID BookID `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"`
	Borrower *UserID `json:"borrower,omitempty"`
}

type User struct {
	ID UserID `json:"id"`
	Name string `json:"name"`
}

