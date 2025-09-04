package internal

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password with its hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		
		// Hash the password before storing
		hashedPassword, err := HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
			return
		}
		
		user := User{Username: req.Username, Password: hashedPassword}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User exists"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Registered"})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		var user User
		// First, find user by username only
		if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		
		// Then verify the password
		if !CheckPasswordHash(req.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not login"})
			return
		}
		c.Header("Authorization", "Bearer "+tokenString)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		var tokenString string
		
		tokenParts := strings.Split(header, " ")
		if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
			tokenString = tokenParts[1]
		}
		
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		
		if err != nil {
			fmt.Println("Error parsing token:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		
		c.Next()
	}
}
