package store

import (
	"github.com/your-username/chatnio/internal/model"
	"github.com/your-username/chatnio/pkg/database"
)

// UserStore 定义了用户数据访问的接口
type UserStore interface {
	// Create 创建新用户
	Create(user *model.User) error
	// GetByUsername 根据用户名查找用户
	GetByUsername(username string) (*model.User, error)
	// GetByID 根据用户ID查找用户
	GetByID(id uint) (*model.User, error)
	// Update 更新用户信息
	Update(user *model.User) error
}

// userStore 实现了 UserStore 接口
type userStore struct {
	db *database.DB // 数据库连接
}

// NewUserStore 创建一个新的 UserStore 实例
func NewUserStore() UserStore {
	return &userStore{db: database.DB}
}

// Create 将新用户保存到数据库
func (s *userStore) Create(user *model.User) error {
	return database.DB.Create(user).Error
}

// GetByUsername 通过用户名查询用户
// 如果用户不存在，返回 error
func (s *userStore) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID 通过用户ID查询用户
// 如果用户不存在，返回 error
func (s *userStore) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (s *userStore) Update(user *model.User) error {
	return database.DB.Save(user).Error
}
