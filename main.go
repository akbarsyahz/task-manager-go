package main

import (
	"net/http"
	"taskManager/api/authn"
	"taskManager/db"
	"taskManager/db/model"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connection()
	r := gin.Default()

	api := r.Group("/api")
	{
		atuhn := api.Group("/authn")
		atuhn.POST("/login", func(ctx *gin.Context) {
			var formLogin authn.Login
			if err := ctx.ShouldBind(&formLogin); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(200, gin.H{
				"massage": "Success",
				"token": "token",
			})
		})

		atuhn.POST("/register", func(ctx *gin.Context) {
			var formLogin authn.Login
			var formUser authn.UserRegister

			if err := ctx.ShouldBind(&formLogin); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := ctx.ShouldBind(&formUser); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			_, errAuthn := authn.CreateUser(
			model.User{
				NameFirst: formUser.NameFirst,
				NameLast: formUser.NameLast,
				Age: formUser.Age,
				DateOfBirth: formUser.DateOfBirth,
				PlaceBirth: formUser.DateOfBirth,
			},
			model.LoginUser{
				Username: formLogin.Username,
				Password: formLogin.Password,
			})

			if errAuthn != nil {
				panic(errAuthn.Error())
			}
			
			ctx.JSON(201, gin.H{
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

	r.Run(":8080")
}
