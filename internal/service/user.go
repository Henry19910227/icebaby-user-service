package service

import (
	"errors"
	"mime/multipart"
	"path"

	"github.com/Henry19910227/icebaby-user-service/internal/model"
	"github.com/Henry19910227/icebaby-user-service/internal/repository"
	"github.com/Henry19910227/icebaby-user-service/internal/validator"
	"github.com/Henry19910227/icebaby-user-service/pkg/upload"
)

type userService struct {
	userRepo repository.UserRepository
	uploader upload.Tool
}

// NewUserService ...
func NewUserService(repo repository.UserRepository, uploader upload.Tool) UserService {
	return &userService{repo, uploader}
}

// GetAll Implement UserService interface
func (us *userService) GetAll() ([]*model.User, error) {
	return us.userRepo.GetAll()
}

func (us *userService) GetUser(email string, password string) (*model.User, error) {
	return us.userRepo.GetUser(email, password)
}

// Get Implement UserService interface
func (us *userService) Get(id int64) (*model.User, error) {
	return us.userRepo.GetByID(id)
}

// Add Implement UserService interface
func (us *userService) Add(validator *validator.UserAddValidator) (int64, error) {
	return us.userRepo.Add(validator.Email, validator.Password, validator.Name, validator.Birthday)
}

// Delete Implement UserService interface
func (us *userService) Delete(validator *validator.UserDeleteValidator) error {
	return us.userRepo.DeleteByID(validator.ID)
}

func (us *userService) UpdateUserinfo(uid int64, name string, birthday string) (*model.User, error) {
	return us.userRepo.UpdateUserinfo(uid, name, birthday)
}

func (us *userService) UpdateEmail(uid int64, email string) (*model.User, error) {
	return us.userRepo.UpdateEmail(uid, email)
}

// UpdatePassword ...
func (us *userService) UpdatePassword(uid int64, oldpwd string, newpwd string) error {
	return nil
}

// UploadImage Implement UserService interface
func (us *userService) UploadImage(id int64, file multipart.File, fileHeader *multipart.FileHeader) error {

	if !us.uploader.CheckUploadImageAllowExt(path.Ext(fileHeader.Filename)) {
		return errors.New("image ext is not allow")
	}

	if !us.uploader.CheckUploadImageMaxSize(file) {
		return errors.New("exceeded maximum file limit")
	}
	newFilename, err := us.uploader.UploadImage(fileHeader)
	if err != nil {
		return err
	}
	if _, err = us.userRepo.UpdateUserImage(id, newFilename); err != nil {
		return err
	}
	return nil
}
