package users

import (
	"context"
	"dfee/api/ent"
	"dfee/api/users/dtos"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceParams for fx
type ServiceParams struct {
	fx.In
	Ent    *ent.Client
	Logger *zap.SugaredLogger
}

// Service for users
type Service struct {
	ent    *ent.Client
	logger *zap.SugaredLogger
}

// NewService creates a new userservice
func NewService(p ServiceParams) *Service {
	return &Service{
		ent:    p.Ent,
		logger: p.Logger,
	}
}

// GetOne gets one user
func (s *Service) GetOne(
	ctx context.Context,
	id int,
) (user *ent.User, err error) {
	user, err = s.ent.User.Get(ctx, id)
	return
}

// GetAll get all the user
func (s *Service) GetAll(
	ctx context.Context,
) (users []*ent.User, err error) {
	users, err = s.ent.User.Query().All(ctx)
	return
}

// Create call to create the user
func (s *Service) Create(
	ctx context.Context,
	dto dtos.AddUser,
) (user *ent.User, err error) {
	err = dto.Validate()
	if err == nil {
		user, err = s.ent.User.Create().SetName(dto.Name).SetAge(21).Save(ctx)
	}
	return
}

// Update updates the user
func (s *Service) Update(
	ctx context.Context,
	id int,
	dto dtos.UpdateUser,
) (user *ent.User, err error) {
	err = dto.Validate()
	if err == nil {
		// todo
		update := s.ent.User.UpdateOneID(id)
		if dto.Age != 0 {
			update = update.SetAge(dto.Age)
		}
		if dto.Name != "" {
			update = update.SetName(dto.Name)
		}
		user, err = update.Save(ctx)
	}
	return
}

// Delete deletes the user
func (s *Service) Delete(
	ctx context.Context,
	id int,
) (err error) {
	err = s.ent.User.DeleteOneID(id).Exec(ctx)
	return
}
