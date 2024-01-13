package routes

import (
	"net/http"

	"polarisai/internal/infrastructure/environment"

	"github.com/gin-gonic/gin"
)

func MakeHealthRoute(r *gin.Engine) {
	version := environment.GetInstance().APP_VERSION
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "version": version})
	})
}
