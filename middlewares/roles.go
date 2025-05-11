package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error" : "No role found"})
			c.Abort()
			return
		}

		userRole, ok := value.(string)
		if !ok || userRole != role {
			c.JSON(http.StatusUnauthorized, gin.H{"error" : "Forbidden: insufficient role"})
			c.Abort()
			return
		}

		c.Next()
	}
}