package jwt

import "time"

// Tool ...
type Tool interface {
	GenerateToken(uid int64) (string, error)
	VerifyToken(token string) error
}

// Setting ...
type Setting interface {
	GetTokenSecret() string
	GetIssuer() string
	GetExpire() time.Duration
}
