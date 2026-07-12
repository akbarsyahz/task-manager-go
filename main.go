package main

import (
	"fmt"
	"taskManager/api/authn"
	"taskManager/api/middleware"
	"taskManager/db"
	"taskManager/docs"
	"taskManager/src"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api
//
//nolint:funlen
func main() {
	db.Connection()
	r := gin.Default()
	r.Use(src.CorsConf())
	docs.SwaggerInfo.BasePath = "/api"
	api := r.Group("/api")
	{
		authnRoute := api.Group("/authn")
		authnRoute.POST("/login", authn.LoginHandler)
		authnRoute.POST("/register", authn.RegistrationHandler)
		// TODO: (Akbar): I think this not proper to put in here, change it to user module itself
		authnRoute.GET("/all-user", authn.GetAllUserHandler)

		taskCore := api.Group("/task")
		taskCore.Use(middleware.AuthMiddleware())
		taskCore.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})

		taskCore.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})

		taskCore.POST("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})

		taskCore.PUT("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})

		taskCore.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	running := r.Run(":8080")
	if running != nil {
		panic(fmt.Errorf("not running"))
	}
}
