package users

import (
	"dfee/api/users/dtos"
	"fmt"
)

// Repository database structure
type Repository struct {
	instances []dtos.User
	// db lib.Database
}

// NewRepository creates a new user repository
func NewRepository(
// db lib.Database
) *Repository {
	return &Repository{
		instances: []dtos.User{
			{ID: 1, Name: "user_1"},
			{ID: 2, Name: "user_2"},
			{ID: 3, Name: "user_3"},
		},
		// db: db,
	}
}

// GetAll gets all users
func (r Repository) GetAll() (users []dtos.User, err error) {
	users = r.instances
	return
}

// Save user
func (r Repository) Save(dto dtos.AddUser) (user dtos.User, err error) {
	user = dtos.User{
		ID:   len(r.instances),
		Name: dto.Name,
	}

	r.instances = append(r.instances, user)
	return
}

// Update updates user
func (r Repository) Update(id int, dto dtos.UpdateUser) (user dtos.User, err error) {
	user, err = r.GetOne(id)
	if err == nil {
		user.Name = dto.Name
	}
	return
}

// GetOne gets ont user
func (r Repository) GetOne(id int) (user dtos.User, err error) {
	for _, v := range r.instances {
		if id == v.ID {
			return v, nil
		}
	}

	return user, fmt.Errorf("user id=%d not found", id)
}

// Delete deletes the row of data
func (r Repository) Delete(id int) error {
	for k, v := range r.instances {
		if id == v.ID {
			r.instances = append(r.instances[:k], r.instances[k+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user id=%d not found", id)
}
