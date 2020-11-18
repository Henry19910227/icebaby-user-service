package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors 解決前端跨域問題
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	}
}
