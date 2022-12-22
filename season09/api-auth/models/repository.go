package models

import "gorm.io/gorm"

type Repository interface {
	CreateUser(acc *User) error
	UpdateUser(acc *User) error
	GetUserByUserName(userName string) (*User, error)
	//GetAccById(id int) (*User, error)
	//DeleteAcc(id int) error
}

type MysqlRepository struct {
	DB *gorm.DB
}

func (r *MysqlRepository) CreateUser(acc *User) error {
	return r.DB.Create(acc).Error
}

func (r *MysqlRepository) UpdateUser(acc *User) error {
	return r.DB.Save(acc).Error
}
func (r *MysqlRepository) GetUserByUserName(userName string) (*User, error) {
	var acc User
	first := r.DB.First(&acc, "user_name = ?", userName)
	if first.Error != nil {
		return nil, first.Error
	}
	return &acc, nil
}
