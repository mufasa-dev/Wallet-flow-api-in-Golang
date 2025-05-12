package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/schemas"
)

// @BasePath /api/v1

// @Sumary Create User
// @Description Create a new user
// @Tags Sigup
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request Body"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /signup [post]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error processing password")
		return
	}

	accountNumer := generateAccountNumer()

	user := schemas.User{
		Name:     request.Name,
		Password: hashedPassword,
		CPF:      request.CPF,
		Account:  accountNumer,
		Wallet:   request.Wallet,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error create opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-user", user)
}

// @BasePath /api/v1

// @Sumary Get Users
// @Description Retrieve a list of registered users (Require authentication)
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} schemas.UserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func ListUserHandler(ctx *gin.Context) {
	users := []schemas.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	sendSuccess(ctx, "list-user", users)
}

// @BasePath /

// @Sumary Get User
// @Description Retrieve a registered user (Require authentication)
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Security BearerAuth
// @Success 200 {object} schemas.UserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [get]
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

	sendSuccess(ctx, "show-user", user)
}

// @BasePath /

// @Sumary Update User
// @Description Update a registered user (Require authentication)
// @Tags Users
// @Accept json
// @Produce json
// @Param request body UpdateUserRequest true "Request Body"
// @Param id query string true "User ID"
// @Security BearerAuth
// @Success 200 {object} schemas.UserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [put]
func UpdateUserHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	if request.Name != "" {
		user.Name = request.Name
	}
	if request.CPF != "" {
		user.CPF = request.CPF
	}
	if request.Account != "" {
		user.Account = request.Account
	}
	if request.Password != "" {
		hashedPassword, err := HashPassword(request.Password)
		if err != nil {
			sendError(ctx, http.StatusInternalServerError, "Error processing password")
			return
		}
		user.Password = hashedPassword
	}

	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, "update-user", request)
}

// @BasePath /

// @Sumary Delete User
// @Description Delete a registered user (Require authentication)
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Security BearerAuth
// @Success 200 {object} schemas.UserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [delete]
func DeleteUserHandler(ctx *gin.Context) {
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

	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting user %s", user.Name))
		return
	}

	sendSuccess(ctx, "deleting-user", user)
}
