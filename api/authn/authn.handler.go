package authn

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler godoc
// @Summary Login user
// @Tags authn
// @Accept json
// @Produce json
// @Param body body LoginDto true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Router /authn/login [post]
// LoginHandler Handle login
func LoginHandler(ctx *gin.Context) {
	var formLogin LoginDto
	if err := ctx.ShouldBind(&formLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultToken, err := Login(formLogin.Username, formLogin.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"token":   resultToken,
	})
}

// RegistrationHandler godoc
// @Summary Register user
// @Tags authn
// @Accept json
// @Produce json
// @Param body body UserRegisterDto true "User credentials"
// @Success 200 {object} map[string]interface{}
// @Router /authn/register [post]
// RegistrationHandler Handle Registration
func RegistrationHandler(ctx *gin.Context) {
	var formUser UserRegisterDto

	if err := ctx.ShouldBind(&formUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, errAuthn := Registration(formUser)
	if errAuthn != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": errAuthn.Error()})
	}

	ctx.JSON(201, gin.H{
		"massage": "Success",
	})
}

// GetAllUserHandler godoc
// @Summary GetAlluser user
// @Tags authn
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /authn/all-user [post]
// RegistrationHandler Handle Registration
func GetAllUserHandler(ctx *gin.Context) {
	users, err := GetAllUser()
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    users,
	})
}
