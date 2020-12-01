package services

import (
	"dfee/api/lib"
	"dfee/api/models"
	"dfee/api/repositories"
)

// UserService service layer
type UserService struct {
	logger         lib.Logger
	userRepository repositories.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger, userRepository repositories.UserRepository) UserService {
	return UserService{
		logger:         logger,
		userRepository: userRepository,
	}
}

// GetOneUser gets one user
func (s UserService) GetOneUser(id int) (user models.User, err error) {
	user, err = s.userRepository.GetOne(id)
	return
}

// GetAllUser get all the user
func (s UserService) GetAllUser() (users []models.User, err error) {
	users, err = s.userRepository.GetAll()
	return
}

// CreateUser call to create the user
func (s UserService) CreateUser(dto models.AddUser) (err error) {
	err = dto.Validate()
	if err == nil {
		_, err = s.userRepository.Save(dto)
	}
	return
}

// UpdateUser updates the user
func (s UserService) UpdateUser(id int, dto models.UpdateUser) (err error) {
	err = dto.Validate()
	if err == nil {
		_, err = s.userRepository.Update(id, dto)
	}
	return
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id int) (err error) {
	err = s.userRepository.Delete(id)
	return
}
