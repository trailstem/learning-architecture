package usecases

import (
	"myapp/entities"
	"myapp/repositories"
	"myapp/usecases/input"
)

type getUser struct {
	repo repositories.UserRepository
}

func NewGetUser(repo repositories.UserRepository) input.GetUserInputBoundary {
	return &getUser{repo: repo}
}

func (u *getUser) GetUser(input input.GetUserInputData) (entities.User, error) {
	user, err := u.repo.GetUser(input.ID)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
