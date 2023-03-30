package controllers

import (
	"myapp/presenters"
	"myapp/repositories"
	"myapp/usecases"
	"myapp/usecases/input"
	"myapp/usecases/output"
)

type UserController struct {
	useCase input.GetUserInputBoundary
}

func NewUserController(repo repositories.UserRepository) *UserController {
	useCase := usecases.NewGetUser(repo)
	return &UserController{useCase: useCase}
}

func (c *UserController) GetUser(id int) (output.GetUserOutputData, error) {
	inputData := input.GetUserInputData{ID: id}
	user, err := c.useCase.GetUser(inputData)
	if err != nil {
		return output.GetUserOutputData{}, err
	}

	presenter := presenters.NewUserPresenter()
	return presenter.Present(user)
}
