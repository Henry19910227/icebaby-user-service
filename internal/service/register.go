package service

import (
	"github.com/Henry19910227/icebaby-user-service/internal/model"
	"github.com/Henry19910227/icebaby-user-service/internal/repository"
	"github.com/Henry19910227/icebaby-user-service/internal/validator"
	"github.com/Henry19910227/icebaby-user-service/utils"
)

// IBRegisterService 註冊服務
type IBRegisterService struct {
	userRepo repository.UserRepository
}

// NewRegisterService 創建註冊服務
func NewRegisterService(repo repository.UserRepository) RegisterService {
	return &IBRegisterService{repo}
}

// Register 註冊用戶
func (service *IBRegisterService) Register(input *validator.Register) (int64, error) {
	uid, err := service.userRepo.GetUserIDByCode(input.Invite)
	if err != nil && len(input.Invite) > 0 {
		return 0, err
	}
	user := &model.User{
		Role:         input.Role,
		Nickname:     input.Nickname,
		Sex:          input.Sex,
		Email:        input.Email,
		InviteCode:   utils.GenerateInviteCode(6),
		InviteUserID: uid,
		Birthday:     input.Birthday,
		AuthType:     input.AythType,
		Identifier:   input.Identifier,
		Password:     input.Password,
	}
	return service.userRepo.Add(user)
}
