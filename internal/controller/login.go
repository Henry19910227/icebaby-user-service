package controller

import (
	"net/http"

	"github.com/Henry19910227/icebaby-user-service/internal/service"
	"github.com/Henry19910227/icebaby-user-service/internal/validator"
	"github.com/Henry19910227/icebaby-user-service/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// LoginController ...
type LoginController struct {
	loginService service.LoginService
	jwtTool      jwt.Tool
}

// NewLoginController ...
func NewLoginController(router *gin.Engine, loginService service.LoginService, tool jwt.Tool) {
	loginController := &LoginController{loginService, tool}
	v1 := router.Group("/icebaby/v1")
	v1.POST("/login", loginController.Login)
}

// Login ...
func (lc *LoginController) Login(c *gin.Context) {
	var validator validator.UserLoginValidator
	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	user, err := lc.loginService.Login(validator.Email, validator.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "帳號或密碼錯誤!"})
		return
	}
	tokenString, err := lc.jwtTool.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "token": tokenString, "data": user, "msg": "login success!"})
}
