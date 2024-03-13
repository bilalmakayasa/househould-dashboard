package middleware

import (
	"household-dashboard/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	Bearer string `header:"Authorization"`
}

func BearerTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token please!!!"})
			return
		}
		bearerToken := strings.Split(h.Bearer, "Bearer")
		tokenValid := strings.TrimSpace(bearerToken[1])

		claims, err := utils.VerifyToken(tokenValid)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
