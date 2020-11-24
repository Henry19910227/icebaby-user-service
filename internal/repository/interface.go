package repository

import "github.com/Henry19910227/icebaby-user-service/internal/model"

// ValidateRepository ...
type ValidateRepository interface {
	ValidateLogin(Identifier string, password string) (int64, error)
	ValidateInviteCode(code string) (int64, error)
}

// UserRepository ...
type UserRepository interface {
	InsertUser(user *model.UserAll) (int64, error)
	GetUserIDByCode(inviteCode string) (int64, error)
	GetUserByID(id int64) (*model.UserAll, error)

	GetAll() ([]*model.UserAll, error)
	DeleteByID(id int64) error
	UpdateUserinfo(uid int64, name string, birthday string) (*model.UserAll, error)
	UpdateEmail(uid int64, email string) (*model.UserAll, error)
	UpdatePassword(uid int64, password string) error
	UpdateUserImage(uid int64, image string) (*model.UserAll, error)
}

// OTPRepository ...
type OTPRepository interface {
	Generate(mobile string) (string, error)
	Validate(code string, secret string) bool
}
