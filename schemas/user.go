package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	CPF      string
	Account  string
	Wallet   float64
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CPF       string    `json:"cpf"`
	Account   string    `json:"account"`
	Wallet    float64   `json:"wallet"`
}
