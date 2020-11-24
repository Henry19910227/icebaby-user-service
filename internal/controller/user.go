package controller

import (
	"net/http"

	"github.com/Henry19910227/icebaby-user-service/internal/service"
	"github.com/Henry19910227/icebaby-user-service/internal/validator"
	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct {
	UserService service.UserService
}

// NewUserController 初始化Controller與創建路由
func NewUserController(router *gin.Engine, userService service.UserService, jwtMidd gin.HandlerFunc) {
	userController := &UserController{
		UserService: userService,
	}
	v1 := router.Group("/icebaby/v1")
	v1.Use(jwtMidd)
	v1.GET("/user", userController.GetAll)
	v1.GET("/user/:id", userController.Get)
	v1.DELETE("/user/:id", userController.DeleteByID)
	v1.PUT("/user/:id/userinfo", userController.UpdateUserinfo)
	v1.PUT("/user/:id/email", userController.UpdateEmail)
	v1.PUT("/user/:id/password", userController.UpdateUserPassword)
	v1.PUT("/user/:id/image", userController.UpdateUserImage)
	v1.StaticFS("/userimage", http.Dir("./storege"))
	v1.GET("/panic", userController.PanicTest)
}

// GetAll 列出所有用戶
func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users, "msg": "success!"})
}

// Get 以 uid 查找單個用戶
func (uc *UserController) Get(c *gin.Context) {
	var validator validator.UserIDValidator
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	user, err := uc.UserService.Get(validator.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "查無此用戶!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user, "msg": "success!"})
}

// DeleteByID 以 uid 刪除用戶
func (uc *UserController) DeleteByID(c *gin.Context) {
	validator := validator.UserDeleteValidator{}
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	if err := uc.UserService.Delete(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "刪除成功!"})
}

// UpdateEmail 更新用戶 Email
func (uc *UserController) UpdateEmail(c *gin.Context) {
	uidValidator := validator.UserIDValidator{}
	if err := c.ShouldBindUri(&uidValidator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	emailValidator := validator.UserEmailValidator{}
	if err := c.ShouldBindJSON(&emailValidator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	user, err := uc.UserService.UpdateEmail(uidValidator.ID, emailValidator.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user, "msg": "update email success!"})
}

// UpdateUserPassword 更新用戶密碼
func (uc *UserController) UpdateUserPassword(c *gin.Context) {
	uidValidator := validator.UserIDValidator{}
	if err := c.ShouldBindUri(&uidValidator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	pwdValidator := validator.UserUpdatePwdValidator{}
	if err := c.ShouldBindJSON(&pwdValidator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	err := uc.UserService.UpdatePassword(uidValidator.ID, pwdValidator.OldPwd, pwdValidator.NewPwd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "update password success!"})
}

// UpdateUserinfo 用戶更新個人資訊
func (uc *UserController) UpdateUserinfo(c *gin.Context) {
	var uidValidator validator.UserIDValidator
	if err := c.ShouldBindUri(&uidValidator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	var validator validator.UserUpdateUserinfoValidator
	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	user, err := uc.UserService.UpdateUserinfo(uidValidator.ID, validator.Name, validator.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user, "msg": "update userinfo success!"})
}

// UpdateUserImage 用戶上傳照片
func (uc *UserController) UpdateUserImage(c *gin.Context) {
	var validator validator.UserIDValidator
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	if err = uc.UserService.UploadImage(validator.ID, file, fileHeader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "upload success!"})
}

// PanicTest 測試 Panic
func (uc *UserController) PanicTest(c *gin.Context) {
	var dict map[string]string
	dict["H"] = "Hello"
}
