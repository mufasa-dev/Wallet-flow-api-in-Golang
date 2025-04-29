package handlers

import (
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetSQLite()
}
