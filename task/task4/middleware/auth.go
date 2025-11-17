package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"github.com/zhaiht/blog_backend/util"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logrus.Warn("Authorization header is required")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization  header is required",
			})
			c.Abort()
			return
		}

		//移除字符串前缀
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		claims := &util.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			logrus.Warnf("Invalid or expired token : %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
