package repositories

import (
	"errors"
	"myapp/entities"
)

// UserRepository インターフェースの具体的な実装
// UserRepository構造体がUserRepositoryインターフェースを実装している
type userRepository struct {
	users []entities.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: []entities.User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		},
	}
}

func (r *userRepository) GetUser(id int) (entities.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return entities.User{}, errors.New("user not found")
}
