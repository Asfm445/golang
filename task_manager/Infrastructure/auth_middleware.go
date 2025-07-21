package infrastructure

import (
	"fmt"
	"strings"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		tokenStr := authParts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token expiration"})
			c.Abort()
			return
		}

		if int64(exp) < time.Now().Unix() {
			c.JSON(401, gin.H{"error": "Token has expired"})
			c.Abort()
			return
		}

		userEmail, ok := claims["email"].(string)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid user ID in token"})
			c.Abort()
			return
		}

		roleClaim, ok := claims["role"].(string)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid role in token"})
			c.Abort()
			return
		}
		if roleClaim == "user" && role == "admin" {
			c.JSON(403, gin.H{"error": "Forbidden: insufficient permissions" + roleClaim})
			c.Abort()
			return
		}
		c.Set("user_email", userEmail)
		c.Next()
	}
}
