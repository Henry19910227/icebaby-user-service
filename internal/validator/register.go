package validator

// Register 驗證用戶註冊
type Register struct {
	AythType   int    `json:"auth_type" binding:"required"`
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8,max=16"`
	Role       int    `json:"role" binding:"required"`
	Nickname   string `json:"nickname" binding:"required,min=1,max=16"`
	Birthday   string `json:"birthday" binding:"required,datetime=2006-01-02"`
	Sex        int    `json:"sex" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Invite     string `json:"invite"`
}

// OTP 驗證OTP請求
type OTP struct {
	Mobile string `json:"mobile" binding:"required"`
}
