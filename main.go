package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	{
		atuhn := r.Group("/authn")
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
	}

	{
		taskCore := r.Group("/task")
		taskCore.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"massage": "Success",
			})
		})
	}

	r.Run(":8080")
}
