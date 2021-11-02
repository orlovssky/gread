package store

// User - An app user
type User struct {
	ID        int    `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname" validate:"required,min=3,max=100"`
	Lastname  string `json:"lastname" db:"lastname" validate:"required,min=3,max=100"`
	Username  string `json:"username" db:"username" validate:"required,min=3,max=30"`
	Password  string `json:"password,omitempty" db:"password" validate:"required,min=8,max=30"`
	Email     string `json:"email" db:"email" validate:"required,email"`
}
