package model

// APIRegisterReq 驗證用戶註冊
type APIRegisterReq struct {
	AythType   int    `json:"auth_type" binding:"required"`
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8,max=16"`
	Role       int    `json:"role" binding:"required"`
	Nickname   string `json:"nickname" binding:"required,min=1,max=16"`
	Birthday   string `json:"birthday" binding:"required,datetime=2006-01-02"`
	Sex        int    `json:"sex" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Invite     string `json:"invite"`
	MobileOTP  string `json:"mobile_otp"`
}

// APIMobileOTPReq 發送OTP請求
type APIMobileOTPReq struct {
	Mobile string `json:"mobile" binding:"required"`
}
