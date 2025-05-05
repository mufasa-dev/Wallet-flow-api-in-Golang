package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/schemas"
)

func StatementHandler(ctx *gin.Context) {
	id, err := GetUserIdFromJWT(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	historic := []schemas.Historic{}

	if err := db.Where("user_id = ?", id).Find(&historic).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "historic not found")
		return
	}

	sendSuccess(ctx, "statement", historic)
}
