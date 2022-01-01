package service

import (
	"com.justin.k8s.api/srv-user/dao"
	"com.justin.k8s.api/srv-user/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *model.User) error
	Login(phone, password string) (*model.User, error)
}

type UserServiceImpl struct {
	userDao dao.UserDao
}

func NewUserService() UserService {
	return &UserServiceImpl{
		userDao: dao.NewUserDao(),
	}
}

func (u *UserServiceImpl) Register(user *model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	return u.userDao.CreateUser(user)
}

func (u *UserServiceImpl) Login(phone, password string) (*model.User, error) {
	user, err := u.userDao.GetUserByPhone(phone)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
