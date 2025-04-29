package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Id       int64
	Name     string
	Password string
	CPF      string
	Account  string
}
