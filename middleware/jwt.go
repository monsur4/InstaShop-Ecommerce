package middleware

import (
	"net/http"
	"strings"
	"time"
	"fmt"
    "os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	// Retrieve JWT Secret Key
	jwtKey = []byte(os.Getenv("SECRET_KEY"))
	if len(jwtKey) == 0 {
		panic("SECRET_KEY is not set in the .env file")
	}

}

type Claims struct {
    UserId uint `json:"user_id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// Generate JWT Token
func GenerateJWT(userId uint, email string, role string) (string, error) {
	claims := &Claims{
	    UserId: userId,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// JWT Middleware
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		fmt.Println(claims)

		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("userId", claims.UserId)
		c.Next()
	}
}

