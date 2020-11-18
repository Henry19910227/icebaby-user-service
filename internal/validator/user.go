package validator

// UserIDValidator ...
type UserIDValidator struct {
	ID int64 `uri:"id" binding:"required,gte=1"`
}

// UserEmailValidator ...
type UserEmailValidator struct {
	Email string `json:"email" binding:"required,email"`
}

// UserLoginValidator ...
type UserLoginValidator struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserDeleteValidator ...
type UserDeleteValidator struct {
	ID int64 `uri:"id" binding:"required,gte=1"`
}

// UserAddValidator ...
type UserAddValidator struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required,datetime=2006-01-02"`
}

// UserUpdatePwdValidator ...
type UserUpdatePwdValidator struct {
	OldPwd string `json:"oldpwd" binding:"required"`
	NewPwd string `json:"newpwd" binding:"required"`
}

// UserUpdateUserinfoValidator ...
type UserUpdateUserinfoValidator struct {
	Name     string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required,datetime=2006-01-02"`
}
