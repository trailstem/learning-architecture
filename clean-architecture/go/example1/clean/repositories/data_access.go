package repositories

import "myapp/entities"

type UserRepository interface {
	GetUser(id int) (entities.User, error)
}
