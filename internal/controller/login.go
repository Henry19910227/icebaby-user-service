package controller

import (
	"net/http"

	"github.com/Henry19910227/icebaby-user-service/internal/model"

	"github.com/Henry19910227/icebaby-user-service/internal/service"
	"github.com/gin-gonic/gin"
)

// LoginController ...
type LoginController struct {
	loginService service.LoginService
}

// NewLoginController ...
func NewLoginController(router *gin.Engine, loginService service.LoginService) {
	loginController := &LoginController{loginService}
	v1 := router.Group("/icebaby/v1")
	v1.POST("/login", loginController.Login)
}

// Login ...
func (lc *LoginController) Login(c *gin.Context) {
	var loginInput model.APILoginReq
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	user, token, err := lc.loginService.Login(loginInput.Identifier, loginInput.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "token": token, "data": user, "msg": "login success!"})
}
