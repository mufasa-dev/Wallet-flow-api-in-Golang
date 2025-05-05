package handlers

import (
	"net/http"
	"time"

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

	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	query := db.Where("user_id = ?", id)

	if startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			sendError(ctx, http.StatusBadRequest, "invalid start_date format, expected YYYY-MM-DD")
			return
		}
		query = query.Where("created_at >= ?", startDate)
	}

	if endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			sendError(ctx, http.StatusBadRequest, "invalid end_date format, expected YYYY-MM-DD")
			return
		}
		// Adiciona 1 dia para incluir o último dia no intervalo
		endDate = endDate.Add(24 * time.Hour)
		query = query.Where("created_at < ?", endDate)
	}

	if err := query.Order("created_at DESC").Find(&historic).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "historic not found")
		return
	}

	sendSuccess(ctx, "statement", historic)
}
