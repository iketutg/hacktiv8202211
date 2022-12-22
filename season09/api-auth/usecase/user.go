package usecase

import (
	"api-auth/models"
)

type UserUseCase interface {
	CreateUser(account *models.User) error
	UpdateUser(account *models.User) error
	GetUserByUserName(userName string) (*models.User, error)
}

type userUseCase struct {
	repository models.Repository
}

func NewUserUseCase(repo models.Repository) UserUseCase {
	return &userUseCase{
		repository: repo,
	}
}

func (uacc *userUseCase) CreateUser(account *models.User) error {
	account.Status = "ACTIVE"
	return uacc.repository.CreateUser(account)
}

func (uacc *userUseCase) UpdateUser(account *models.User) error {
	return uacc.repository.UpdateUser(account)
}

func (uacc *userUseCase) GetUserByUserName(userName string) (*models.User, error) {
	return uacc.repository.GetUserByUserName(userName)
}
