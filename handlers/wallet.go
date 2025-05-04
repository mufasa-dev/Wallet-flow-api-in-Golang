package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/schemas"
)

func WithdrawHandler(ctx *gin.Context) {
	request := DepositWithDrawRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := GetUserIdFromJWT(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		logger.Errorf("error finding user by id: %v", id)
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	if user.Wallet < request.Amount {
		sendError(ctx, http.StatusInternalServerError, "not enough money in the wallet")
		return
	}

	user.Wallet = user.Wallet - request.Amount

	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, "withdraw", request)
}

func DepositHandler(ctx *gin.Context) {
	request := DepositWithDrawRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := GetUserIdFromJWT(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		logger.Errorf("error finding user by id: %v", id)
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	user.Wallet = user.Wallet + request.Amount

	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, "deposit", request)
}
