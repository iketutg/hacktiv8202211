package usecase

import (
	"api-account/models"
	"github.com/google/uuid"
)

type AccountUseCase interface {
	CreateAccount(account *models.Acccount) error
	UpdateAccount(account *models.Acccount) error
	//GetAccountByPhoneNumber(phoneNumber string) (*models.User, error)
	UpdateStatusAccount(phoneNumber string) (*models.Acccount, error)
}

type accountUseCase struct {
	repository models.Repository
}

func NewAccountUseCase(repo models.Repository) AccountUseCase {
	return &accountUseCase{
		repository: repo,
	}
}

/*
	{
	     "UserName":"IKetutG",
	    "Password":"123456",
	    "FullName":"IKetut Gnawan",
		"Email":         "iketutg@gmail.com",
		"PhoneNumber":   "08121323410"
	}
*/
func (uacc *accountUseCase) CreateAccount(account *models.Acccount) error {
	//generate Accounter
	//Set status OPEN
	uuidstr := uuid.New()
	account.AccountNumber = uuidstr.String()
	account.Status = "OPEN"
	return uacc.repository.CreateAcc(account)
}

func (uacc *accountUseCase) UpdateAccount(account *models.Acccount) error {
	return uacc.repository.UpdateAcc(account)
}

func (uacc *accountUseCase) UpdateStatusAccount(phoneNumber string) (*models.Acccount, error) {
	account, err := uacc.repository.GetAccByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}
	account.Status = "ACTIVE"
	err = uacc.repository.UpdateAcc(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}
