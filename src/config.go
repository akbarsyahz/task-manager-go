package src

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsConf for config thee CORS
func CorsConf() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})
		c.Next()
	}
}
