package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/schemas"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = os.Getenv("JWT_SECRET")

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var storedUser schemas.User
	result := db.Where("Name = ?", req.Username).Find(&storedUser)
	if result.Error != nil || storedUser.ID == 0 {
		sendError(c, http.StatusNotFound, "user not found")
		return
	}

	if !CheckPasswordHash(req.Password, storedUser.Password) {
		sendError(c, http.StatusNotFound, "Incorrect password")
		return
	}

	token, err := GenetareJWT(req.Username)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	sendSuccess(c, "sigin", gin.H{"token": token})
}

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

	user := schemas.User{
		Name:     request.Name,
		Password: hashedPassword,
		CPF:      request.CPF,
		Wallet:   request.Wallet,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error create opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-user", user)
}

func GenetareJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString((jwtSecret))
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
