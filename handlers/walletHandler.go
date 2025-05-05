package handlers

import (
	"fmt"
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

	tx := db.Begin()

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

	user.Wallet -= request.Amount

	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	historic := schemas.Historic{
		Action:  "withdraw",
		Comment: "",
		UserId:  id,
	}

	if err := db.Create(&historic).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error saving historic")
		return
	}

	if err := tx.Commit().Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "transaction commit failed")
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

	tx := db.Begin()

	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		logger.Errorf("error finding user by id: %v", id)
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	user.Wallet += request.Amount

	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	historic := schemas.Historic{
		Action:  "deposit",
		Comment: "",
		UserId:  id,
	}

	if err := db.Create(&historic).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error saving historic")
		return
	}

	if err := tx.Commit().Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "transaction commit failed")
		return
	}

	sendSuccess(ctx, "deposit", request)
}

func TransferHandler(ctx *gin.Context) {
	request := TransferRequest{}

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

	recipient := schemas.User{}
	if err := db.Where("CPF = ?", request.RecipientCPF).First(&recipient).Error; err != nil {
		logger.Errorf("error finding user recipient by cpf: %v", request.RecipientCPF)
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	if user.ID == recipient.ID {
		sendError(ctx, http.StatusInternalServerError, "Its not allow to make a transfer to yourself")
		return
	}

	tx := db.Begin()

	user.Wallet -= request.Amount
	recipient.Wallet += request.Amount

	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user sender: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	if err := db.Save(&recipient).Error; err != nil {
		logger.Errorf("error updating user recipient: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	historic := schemas.Historic{
		Action:  "transfer",
		Comment: fmt.Sprintf("transfer R$%v to %v", request.Amount, recipient.Name),
		UserId:  id,
	}

	if err := db.Create(&historic).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error saving historic")
		return
	}

	if err := tx.Commit().Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "transaction commit failed")
		return
	}

	sendSuccess(ctx, "withdraw", request)
}
