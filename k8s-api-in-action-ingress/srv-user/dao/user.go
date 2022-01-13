package dao

import (
	"com.justin.k8s.api/common/databases"
	"com.justin.k8s.api/srv-user/model"
)

type UserDao interface {
	CreateUser(user *model.User) error
	GetUsers(pageSize, pageNum int) ([]model.User, int, error)
	UpdateUser(id int, user *model.User) error
	DeleteUser(id int) error
	GetUserByPhone(phone string) (*model.User, error)
}

type UserDaoImpl struct{}

func NewUserDao() UserDao {
	return &UserDaoImpl{}
}

func (u *UserDaoImpl) CreateUser(user *model.User) error {
	return databases.DB.Create(user).Error
}

func (u *UserDaoImpl) GetUsers(pageSize, pageNum int) ([]model.User, int, error) {
	var users []model.User
	var total int

	err := databases.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (u *UserDaoImpl) UpdateUser(id int, user *model.User) error {
	return databases.DB.Model(&model.User{}).Where("id = ?", id).Update(user).Error
}

func (u *UserDaoImpl) DeleteUser(id int) error {
	return databases.DB.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u *UserDaoImpl) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	err := databases.DB.Where("phone = ?", phone).First(&user).Error
	return &user, err
}
