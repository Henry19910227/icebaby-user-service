package repository

import "github.com/Henry19910227/icebaby-user-service/internal/model"

// UserRepository ...
type UserRepository interface {
	GetAll() ([]*model.User, error)
	GetUser(email string, password string) (*model.User, error)
	GetByID(id int64) (*model.User, error)
	Add(email string, password string, name string, birthday string) (int64, error)
	DeleteByID(id int64) error
	UpdateUserinfo(uid int64, name string, birthday string) (*model.User, error)
	UpdateEmail(uid int64, email string) (*model.User, error)
	UpdatePassword(uid int64, password string) error
	UpdateUserImage(uid int64, image string) (*model.User, error)
}
