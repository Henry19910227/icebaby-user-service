package repository

import (
	"github.com/Henry19910227/icebaby-user-service/pkg/otp"
)

// ICOTPReopsitory ...
type ICOTPReopsitory struct {
	otpTool otp.Tool
}

// NewOTPReopsitory ...
func NewOTPReopsitory(otpTool otp.Tool) OTPRepository {
	return &ICOTPReopsitory{otpTool}
}

// Generate 生成 OTP 密碼
func (repo *ICOTPReopsitory) Generate(mobile string) (string, error) {
	return repo.otpTool.Generate(mobile)
}

// Validate 驗證 OTP 密碼是否有效
func (repo *ICOTPReopsitory) Validate(code string, secret string) bool {
	return repo.otpTool.Validate(code, secret)
}
