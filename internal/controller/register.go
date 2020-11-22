package controller

import (
	"net/http"

	"github.com/Henry19910227/icebaby-user-service/internal/service"
	"github.com/Henry19910227/icebaby-user-service/internal/validator"
	"github.com/gin-gonic/gin"
)

// ICRegisterController ...
type ICRegisterController struct {
	registerService service.RegisterService
}

// NewRegisterController ...
func NewRegisterController(router *gin.Engine, registerService service.RegisterService) {
	registerController := &ICRegisterController{registerService}
	v1 := router.Group("/icebaby/v1")
	v1.POST("/register", registerController.Register)
	v1.POST("/register/send_mobile_otp", registerController.SendMobileOTP)
}

// Register ...
func (rc *ICRegisterController) Register(c *gin.Context) {
	var input validator.Register
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	uid, err := rc.registerService.Register(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": uid, "msg": "register success!"})
}

// SendMobileOTP 生成 Phone OTP
func (rc *ICRegisterController) SendMobileOTP(c *gin.Context) {
	var input validator.OTP
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	code, err := rc.registerService.SendMobileOTP(input.Mobile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": code, "msg": "驗證簡訊已寄出!"})
}
