package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recover ...
func Recover() func(c *gin.Context, recovered interface{}) {
	return func(c *gin.Context, recovered interface{}) {
		// 獲取自定義的panic
		if str, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "data": nil, "msg": str})
			c.Abort() //終止後續調用
			return
		}
		// 獲取系統的panic
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "data": nil, "msg": "發生不知名錯誤!"})
		c.Abort() //終止後續調用
	}
}
