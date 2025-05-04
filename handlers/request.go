package handlers

import (
	"fmt"

	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/utils"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}

func errParamIsInvalid(name string) error {
	return fmt.Errorf("param %s is invalid", name)
}

// Create user
type CreateUserRequest struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	CPF      string  `json:"cpf"`
	Account  string  `json:"account"`
	Wallet   float64 `json:"wallet"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.CPF == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	if r.CPF == "" {
		return errParamIsRequired("CPF", "string")
	}
	if !utils.ValidateCPF(r.CPF) {
		return errParamIsInvalid("CPF")
	}
	return nil
}

// Update user
type UpdateUserRequest struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	CPF      string  `json:"cpf"`
	Account  string  `json:"account"`
	Wallet   float64 `json:"wallet"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.CPF != "" && !utils.ValidateCPF(r.CPF) {
		return errParamIsInvalid("CPF")
	}
	if r.Name != "" || r.CPF != "" || r.Password != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

// Deposit or withdraw
type DepositWithDrawRequest struct {
	Amount float64 `json:"amount"`
}

func (r *DepositWithDrawRequest) Validate() error {
	if r.Amount <= 0 {
		return errParamIsInvalid("Amount")
	}
	return nil
}

// Transfer
type TransferRequest struct {
	RecipientCPF string  `json:"recipient_cpf"`
	Amount       float64 `json:"amount"`
}

func (r *TransferRequest) Validate() error {
	if r.RecipientCPF == "" {
		return errParamIsRequired("recipient_cpf", "string")
	}
	if !utils.ValidateCPF(r.RecipientCPF) {
		return errParamIsInvalid("recipient_cpf")
	}
	if r.Amount <= 0 {
		return errParamIsInvalid("amount")
	}
	return nil
}
