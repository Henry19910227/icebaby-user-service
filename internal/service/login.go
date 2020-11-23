package service

import (
	"github.com/Henry19910227/icebaby-user-service/internal/model"
	"github.com/Henry19910227/icebaby-user-service/internal/repository"
)

type loginService struct {
	userRepo     repository.UserRepository
	validateRepo repository.ValidateRepository
}

// NewLoginService ...
func NewLoginService(userRepo repository.UserRepository, validateRepo repository.ValidateRepository) LoginService {
	return &loginService{userRepo, validateRepo}
}

// Login ...
func (service *loginService) Login(mobile string, password string) (*model.APILoginRes, error) {
	uid, err := service.validateRepo.Validate(mobile, password)
	if err != nil {
		return nil, err
	}
	user, err := service.userRepo.GetByID(uid)
	if err != nil {
		return nil, err
	}
	res := &model.APILoginRes{
		ID:       user.ID,
		Role:     user.Role,
		Status:   user.Status,
		Nickname: user.Nickname,
	}
	return res, nil
}
