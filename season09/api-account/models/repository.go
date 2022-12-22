package models

import "gorm.io/gorm"

type Repository interface {
	CreateAcc(acc *Acccount) error
	UpdateAcc(acc *Acccount) error
	GetAccByPhoneNumber(phoneNumber string) (*Acccount, error)
	//GetAccById(id int) (*Acccount, error)
	//DeleteAcc(id int) error
}

type MysqlRepository struct {
	DB *gorm.DB
}

func (r *MysqlRepository) CreateAcc(acc *Acccount) error {
	return r.DB.Create(acc).Error
}

func (r *MysqlRepository) UpdateAcc(acc *Acccount) error {
	return r.DB.Save(acc).Error
}
func (r *MysqlRepository) GetAccByPhoneNumber(phoneNumber string) (*Acccount, error) {
	var acc Acccount
	first := r.DB.First(&acc, "phone_number = ?", phoneNumber)
	if first.Error != nil {
		return nil, first.Error
	}
	return &acc, nil
}
