package middleware

import (
	"net/http"

	"github.com/Henry19910227/icebaby-user-service/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// JWT ...
func JWT(jwtTool jwt.Tool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "token not null"})
			c.Abort()
			return
		}
		if err := jwtTool.VerifyToken(token); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
			c.Abort()
			return
		}
	}
}
