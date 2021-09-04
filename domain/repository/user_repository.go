package repository

import (
	"github.com/wangjiandev/user/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// InitTable 初始化数据表
	InitTable() error
	// FindUserByName 根据用户名查询用户
	FindUserByName(name string) (*model.User, error)
	// FindUserById 根据ID查询用户
	FindUserById(id int64) (*model.User, error)
	// CreateUser 创建用户
	CreateUser(user *model.User) (int64, error)
	// DeleteUserById 根据id删除用户
	DeleteUserById(id int64) error
	// UpdateUser 更新用户
	UpdateUser(user *model.User) error
	// FindAll 查询全部用户
	FindAll() ([]model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) InitTable() error {
	return u.db.AutoMigrate(&model.User{})
}

func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	return user, u.db.Where("user_name = ?", name).Find(user).Error
}

func (u *UserRepository) FindUserById(id int64) (*model.User, error) {
	user := &model.User{}
	return user, u.db.First(user, id).Error
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.db.Create(user).Error
}

func (u *UserRepository) DeleteUserById(id int64) error {
	return u.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.db.Save(user).Error
}

func (u *UserRepository) FindAll() (users []model.User, err error) {
	return users, u.db.Find(&users).Error
}

// NewUserRepository 创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
