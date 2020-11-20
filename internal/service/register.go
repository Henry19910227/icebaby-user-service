package service

import (
	"github.com/Henry19910227/icebaby-user-service/internal/repository"
	"github.com/Henry19910227/icebaby-user-service/internal/validator"
)

// ICERegisterService 註冊服務
type ICERegisterService struct {
	userRepo repository.UserRepository
}

// NewRegisterService 創建註冊服務
func NewRegisterService(repo repository.UserRepository) RegisterService {
	return &ICERegisterService{repo}
}

// Register 註冊用戶
func (service *ICERegisterService) Register(user validator.Register) (int64, error) {
	return 0, nil
}
