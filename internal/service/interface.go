package service

import (
	"mime/multipart"

	"github.com/Henry19910227/icebaby-user-service/internal/model"
	"github.com/Henry19910227/icebaby-user-service/internal/validator"
)

// LoginService ...
type LoginService interface {
	Login(mobile string, password string) (*model.APILoginRes, string, error)
}

// RegisterService ...
type RegisterService interface {
	Register(user *model.APIRegisterReq) (int64, error)
	SendMobileOTP(mobile string) (string, error)
	VerifyMobileOTP(code string, mobile string) bool
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
