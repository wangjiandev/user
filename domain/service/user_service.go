package service

import (
	"github.com/pkg/errors"
	"github.com/wangjiandev/user/domain/model"
	"github.com/wangjiandev/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(id int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(name string) (*model.User, error)
	CheckPwd(userName string, pwd string) (bool, error)
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

type UserService struct {
	userRepository repository.IUserRepository
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password string, hashPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false, errors.New("用户名或密码错误")
	}
	return true, nil
}

func (u *UserService) AddUser(user *model.User) (int64, error) {
	password, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(password)
	return u.userRepository.CreateUser(user)
}

func (u *UserService) DeleteUser(id int64) error {
	return u.userRepository.DeleteUserById(id)
}

func (u *UserService) UpdateUser(user *model.User, isChangePwd bool) error {
	if isChangePwd {
		password, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(password)
	}
	return u.userRepository.UpdateUser(user)
}

func (u *UserService) FindUserByName(name string) (*model.User, error) {
	return u.userRepository.FindUserByName(name)
}

func (u *UserService) CheckPwd(userName string, pwd string) (bool, error) {
	user, err := u.userRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return VerifyPassword(pwd, user.HashPassword)
}
