package repositories

import (
	"dfee/api/models"
	"fmt"
)

// UserRepository database structure
type UserRepository struct {
	instances []models.User
	// db lib.Database
}

// NewUserRepository creates a new user repository
func NewUserRepository(
// db lib.Database
) UserRepository {
	return UserRepository{
		instances: []models.User{
			{ID: 1, Name: "user_1"},
			{ID: 2, Name: "user_2"},
			{ID: 3, Name: "user_3"},
		},
		// db: db,
	}
}

// GetAll gets all users
func (r UserRepository) GetAll() (users []models.User, err error) {
	users = r.instances
	return
}

// Save user
func (r UserRepository) Save(dto models.AddUser) (user models.User, err error) {
	user = models.User{
		ID:   len(r.instances),
		Name: dto.Name,
	}

	r.instances = append(r.instances, user)
	return
}

// Update updates user
func (r UserRepository) Update(id int, dto models.UpdateUser) (user models.User, err error) {
	user, err = r.GetOne(id)
	if err == nil {
		user.Name = dto.Name
	}
	return
}

// GetOne gets ont user
func (r UserRepository) GetOne(id int) (user models.User, err error) {
	for _, v := range r.instances {
		if id == v.ID {
			return v, nil
		}
	}

	return user, fmt.Errorf("user id=%d not found", id)
}

// Delete deletes the row of data
func (r UserRepository) Delete(id int) error {
	for k, v := range r.instances {
		if id == v.ID {
			r.instances = append(r.instances[:k], r.instances[k+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user id=%d not found", id)
}
