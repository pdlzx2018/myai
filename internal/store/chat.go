package store

import (
	"github.com/your-username/chatnio/internal/model"
	"github.com/your-username/chatnio/pkg/database"
)

// ChatStore 定义了聊天记录数据访问的接口
type ChatStore interface {
	// Create 创建新的聊天记录
	Create(chat *model.Chat) error
	// GetByUserID 获取指定用户的聊天记录，支持分页
	GetByUserID(userID uint, limit, offset int) ([]model.Chat, error)
	// GetByID 根据ID获取单条聊天记录
	GetByID(id uint) (*model.Chat, error)
}

// chatStore 实现了 ChatStore 接口
type chatStore struct {
	db *database.DB // 数据库连接
}

// NewChatStore 创建一个新的 ChatStore 实例
func NewChatStore() ChatStore {
	return &chatStore{db: database.DB}
}

// Create 保存新的聊天记录到数据库
func (s *chatStore) Create(chat *model.Chat) error {
	return database.DB.Create(chat).Error
}

// GetByUserID 获取用户的聊天历史记录
// userID: 用户ID
// limit: 每页记录数
// offset: 偏移量（跳过多少条记录）
func (s *chatStore) GetByUserID(userID uint, limit, offset int) ([]model.Chat, error) {
	var chats []model.Chat
	err := database.DB.Where("user_id = ?", userID).
		Order("created_at DESC"). // 按创建时间倒序排序
		Limit(limit).             // 限制返回数量
		Offset(offset).           // 设置偏移量
		Find(&chats).Error
	return chats, err
}

// GetByID 通过ID获取单条聊天记录
func (s *chatStore) GetByID(id uint) (*model.Chat, error) {
	var chat model.Chat
	if err := database.DB.First(&chat, id).Error; err != nil {
		return nil, err
	}
	return &chat, nil
}
