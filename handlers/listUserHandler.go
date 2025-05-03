package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/schemas"
)

func ListUserHandler(ctx *gin.Context) {
	users := []schemas.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	sendSuccess(ctx, "list-user", users)
}

func ShowUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	user := schemas.User{}

	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id %s not found", id))
		return
	}

	if err := db.First(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting user %s", user.Name))
		return
	}

	sendSuccess(ctx, "show-user", user)
}
