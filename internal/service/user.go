package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/pdlzx2018/myai/internal/model"
	"github.com/pdlzx2018/myai/internal/store"
	"github.com/pdlzx2018/myai/pkg/utils"
)

// UserService 定义了用户相关的业务逻辑接口
type UserService interface {
	// Register 用户注册
	Register(username, password, email string) error
	// Login 用户登录
	Login(username, password string) (string, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(id uint) (*model.User, error)
	// UpdateUserInfo 更新用户信息
	UpdateUserInfo(user *model.User) error
}

// userService 实现了 UserService 接口
type userService struct {
	userStore store.UserStore // 用户数据访问对象
}

// NewUserService 创建一个新的 UserService 实例
func NewUserService() UserService {
	return &userService{
		userStore: store.NewUserStore(),
	}
}

// Register 处理用户注册逻辑
// 1. 检查用户名是否已存在
// 2. 对密码进行加密
// 3. 创建新用户
func (s *userService) Register(username, password, email string) error {
	// 检查用户名是否已存在
	if _, err := s.userStore.GetByUsername(username); err == nil {
		return errors.New("用户名已存在")
	}

	// 使用 bcrypt 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 创建新用户
	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	return s.userStore.Create(user)
}

// Login 处理用户登录逻辑
// 1. 验证用户名和密码
// 2. 生成 JWT token
func (s *userService) Login(username, password string) (string, error) {
	// 获取用户信息
	user, err := s.userStore.GetByUsername(username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 生成 JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserInfo 获取用户信息
func (s *userService) GetUserInfo(id uint) (*model.User, error) {
	return s.userStore.GetByID(id)
}

// UpdateUserInfo 更新用户信息
func (s *userService) UpdateUserInfo(user *model.User) error {
	return s.userStore.Update(user)
}
