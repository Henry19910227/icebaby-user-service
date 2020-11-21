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
	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": uid, "msg": "register success!"})
}
