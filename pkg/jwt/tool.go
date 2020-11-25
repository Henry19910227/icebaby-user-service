package jwt

import (
	"errors"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GPJWTTool ...
type GPJWTTool struct {
	setting Setting
}

// NewJWTTool ...
func NewJWTTool(setting Setting) *GPJWTTool {
	return &GPJWTTool{setting}
}

// GenerateToken ...
func (t *GPJWTTool) GenerateToken(uid int64) (string, error) {
	claims := jwt.MapClaims{"sub": strconv.Itoa(int(uid)), "exp": time.Now().Add(t.setting.GetExpire()).Unix()}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.setting.GetTokenSecret()))
	return token, err
}

// VerifyToken ...
func (t *GPJWTTool) VerifyToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.setting.GetTokenSecret()), nil
	})
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return errors.New("Timeout")
		default:
			return err
		}
	}
	return nil
}
