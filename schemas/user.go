package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int64
	Name     string
	Password string
	CPF      string
	Account  string
}
