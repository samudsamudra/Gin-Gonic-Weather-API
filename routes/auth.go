package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"weather-api/config"
)

var jwtSecret = []byte("secret_key")

func Register(c *gin.Context) {
	type RegisterInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword := hashPassword(input.Password)
	_, err := config.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", input.Username, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var hashedPassword string
	var userID int
	err := config.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", input.Username).Scan(&userID, &hashedPassword)
	if err != nil || hashPassword(input.Password) != hashedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := generateJWT(userID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func generateJWT(userID int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}
