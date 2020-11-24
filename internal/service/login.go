package service

import (
	"github.com/Henry19910227/icebaby-user-service/internal/model"
	"github.com/Henry19910227/icebaby-user-service/internal/repository"
	"github.com/Henry19910227/icebaby-user-service/pkg/jwt"
)

type loginService struct {
	userRepo     repository.UserRepository
	validateRepo repository.ValidateRepository
	jwtTool      jwt.Tool
}

// NewLoginService ...
func NewLoginService(userRepo repository.UserRepository, validateRepo repository.ValidateRepository, jwtTool jwt.Tool) LoginService {
	return &loginService{userRepo, validateRepo, jwtTool}
}

// Login ...
func (service *loginService) Login(mobile string, password string) (*model.APILoginRes, string, error) {
	uid, err := service.validateRepo.ValidateLogin(mobile, password)
	if err != nil {
		return nil, "", err
	}
	user, err := service.userRepo.GetByID(uid)
	if err != nil {
		return nil, "", err
	}
	res := &model.APILoginRes{
		ID:       user.ID,
		Role:     user.Role,
		Status:   user.Status,
		Nickname: user.Nickname,
	}
	token, err := service.jwtTool.GenerateToken(uid)
	if err != nil {
		return nil, "", err
	}
	return res, token, nil
}
