package main

import (
	"taskManager/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connection()
	r := gin.Default()

	api := r.Group("/api")
	{
		atuhn := api.Group("/authn")
		atuhn.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})
		atuhn.POST("/register", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})

		atuhn.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})

		taskCore := api.Group("/task")
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

		taskCore.PATCH("/", func(ctx *gin.Context) {
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

	r.Run(":8080")
}
