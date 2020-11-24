package service

import (
	"github.com/Henry19910227/icebaby-user-service/internal/model"
	"github.com/Henry19910227/icebaby-user-service/internal/repository"
	"github.com/Henry19910227/icebaby-user-service/pkg/otp"
	"github.com/Henry19910227/icebaby-user-service/utils"
)

// IBRegisterService 註冊服務
type IBRegisterService struct {
	userRepo     repository.UserRepository
	validateRepo repository.ValidateRepository
	otpTool      otp.Tool
}

// NewRegisterService 創建註冊服務
func NewRegisterService(userRepo repository.UserRepository, validateRepo repository.ValidateRepository, otpTool otp.Tool) RegisterService {
	return &IBRegisterService{userRepo, validateRepo, otpTool}
}

// Register 註冊用戶
func (service *IBRegisterService) Register(input *model.APIRegisterReq) (int64, error) {
	uid, err := service.validateRepo.ValidateInviteCode(input.Invite)
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
		Status:       1,
	}
	return service.userRepo.InsertUser(user)
}

// SendMobileOTP ...
func (service *IBRegisterService) SendMobileOTP(mobile string) (string, error) {
	return service.otpTool.Generate(mobile)
}

// VerifyMobileOTP ...
func (service *IBRegisterService) VerifyMobileOTP(code string, mobile string) bool {
	return service.otpTool.Validate(code, mobile)
}
