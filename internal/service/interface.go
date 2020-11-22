package service

import (
	"mime/multipart"

	"github.com/Henry19910227/icebaby-user-service/internal/model"
	"github.com/Henry19910227/icebaby-user-service/internal/validator"
)

// LoginService ...
type LoginService interface {
	Login(email string, password string) (*model.User, error)
}

//  AythType   int    `json:"auth_type" binding:"required"`
// 	Identifier string `json:"identifier" binding:"required"`
// 	Password   string `json:"password" binding:"required,min=8,max=16"`
// 	Role       int    `json:"role" binding:"required"`
// 	Nickname   string `json:"nickname" binding:"required,min=1,max=16"`
// 	Birthday   string `json:"birthday" binding:"required,datetime=2006-01-02"`
// 	Email      string `json:"email" binding:"required,email"`
// 	InviteCode string `json:"invite_code"`

// RegisterService ...
type RegisterService interface {
	Register(user *validator.Register) (int64, error)
	SendMobileOTP(mobile string) (string, error)
}

// UserService ...
type UserService interface {
	GetAll() ([]*model.User, error)
	GetUser(email string, password string) (*model.User, error)
	Get(id int64) (*model.User, error)
	Add(validator *validator.UserAddValidator) (int64, error)
	Delete(validator *validator.UserDeleteValidator) error
	UpdateUserinfo(uid int64, name string, birthday string) (*model.User, error)
	UpdateEmail(uid int64, email string) (*model.User, error)
	UpdatePassword(uid int64, oldpwd string, newpwd string) error
	UploadImage(id int64, file multipart.File, fileHeader *multipart.FileHeader) error
}
