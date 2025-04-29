package handlers

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
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
		return errParamIsRequired("name", "string")
	}
	if r.CPF == "" {
		return errParamIsRequired("name", "string")
	}
	return nil
}
