package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/schemas"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// @BasePath /

// @Sumary Create User
// @Description Create a new user
// @Tags Sigin
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Request Body"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /signin [post]
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

	token, err := GenetareJWT(storedUser.ID, req.Username)
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	sendSuccess(c, "sigin", gin.H{"token": token})
}

func GenetareJWT(id uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  id,
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

func generateAccountNumer() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%08d", rand.Intn(100000000))
}

func GetUserIdFromJWT(c *gin.Context) (uint, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, http.ErrNoCookie
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["user_id"].(float64)
		return uint(userId), nil
	}

	return 0, jwt.ErrSignatureInvalid
}
