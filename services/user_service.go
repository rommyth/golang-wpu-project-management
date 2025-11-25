package services

import (
	"errors"
	"project-management/models"
	"project-management/repositories"
	"project-management/utils"

	"github.com/google/uuid"
)

type UserServices interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByPublicID(publicID string) (*models.User, error)
	GetAllPagination(filter, sort string, limit, offset int) ([]models.User, int64, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserServices {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	// Cek email apakah sudah terdaftar
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("email already registerd")
	}

	// hasing password
	hashsed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashsed

	// set role
	user.Role = "user"
	user.PublicID = uuid.New()

	// simpan user
	return s.repo.Create(user)
}

func (s *userService) Login(email string, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return user, errors.New("Invalid Credential")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("Invalid Credential")
	}

	return user, nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetByPublicID(publicID string) (*models.User, error) {
	return s.repo.FindByPublicID(publicID)
}

func (s *userService) GetAllPagination(filter, sort string, limit, offset int) ([]models.User, int64, error) {
	return s.repo.FindAllPagination(filter, sort, limit, offset)
}

func (s *userService) Update(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
