package users

import (
	"dfee/api/users/dtos"

	"go.uber.org/zap"
)

// Service service layer
type Service struct {
	logger         *zap.SugaredLogger
	userRepository *Repository
}

// NewService creates a new userservice
func NewService(logger *zap.SugaredLogger, userRepository *Repository) *Service {
	return &Service{
		logger:         logger,
		userRepository: userRepository,
	}
}

// GetOneUser gets one user
func (s *Service) GetOneUser(id int) (user dtos.User, err error) {
	user, err = s.userRepository.GetOne(id)
	return
}

// GetAllUser get all the user
func (s *Service) GetAllUser() (users []dtos.User, err error) {
	users, err = s.userRepository.GetAll()
	return
}

// CreateUser call to create the user
func (s *Service) CreateUser(dto dtos.AddUser) (err error) {
	err = dto.Validate()
	if err == nil {
		_, err = s.userRepository.Save(dto)
	}
	return
}

// UpdateUser updates the user
func (s *Service) UpdateUser(id int, dto dtos.UpdateUser) (err error) {
	err = dto.Validate()
	if err == nil {
		_, err = s.userRepository.Update(id, dto)
	}
	return
}

// DeleteUser deletes the user
func (s *Service) DeleteUser(id int) (err error) {
	err = s.userRepository.Delete(id)
	return
}
