package controller

import (
	"test_mnc/model/dto"
	"test_mnc/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUC usecase.UserUseCase
	authUC usecase.AuthUseCase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload dto.AuthRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	authResponse, err := a.authUC.Login(payload)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"message": "successfully login",
		"data":    authResponse,
	}

	c.JSON(201, response)
}

func (a *AuthController) registerHandler(c *gin.Context) {
	var payload dto.AuthRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.Register(payload)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"message": "successfully register",
	}

	c.JSON(201, response)
}

func (a *AuthController) logoutHandler(c *gin.Context) {
	email := "rizki@gmail.com"
	err := a.authUC.Logout(email) // Ganti email dengan cara Anda mengidentifikasi pengguna.

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Gagal logout",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Berhasil logout",
	})
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/login", a.loginHandler)
	a.rg.POST("/auth/register", a.registerHandler)
	a.rg.POST("/auth/logout", a.logoutHandler)
}

func NewAuthController(userUC usecase.UserUseCase, authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		userUC: userUC,
		authUC: authUC,
		rg:     rg,
	}
}
