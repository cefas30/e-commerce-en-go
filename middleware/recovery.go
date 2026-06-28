package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {

	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error interno del servidor",
		})

	})

}
