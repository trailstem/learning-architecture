package presenters

import (
	"myapp/entities"
	"myapp/usecases/output"
)

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) Present(user entities.User) (output.GetUserOutputData, error) {
	outputData := output.GetUserOutputData{
		ID:   user.ID,
		Name: user.Name,
	}
	return outputData, nil
}
