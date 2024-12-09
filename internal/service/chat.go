package service

import (
	"github.com/your-username/chatnio/internal/model"
	"github.com/your-username/chatnio/internal/store"
)

// ChatService 定义了聊天相关的业务逻辑接口
type ChatService interface {
	// SendMessage 发送消息并获取AI响应
	SendMessage(userID uint, message, modelName string) (*model.Chat, error)
	// GetChatHistory 获取聊天历史记录
	GetChatHistory(userID uint, page, pageSize int) ([]model.Chat, error)
}

// chatService 实现了 ChatService 接口
type chatService struct {
	chatStore store.ChatStore // 聊天记录数据访问对象
}

// NewChatService 创建一个新的 ChatService 实例
func NewChatService() ChatService {
	return &chatService{
		chatStore: store.NewChatStore(),
	}
}

// SendMessage 处理发送消息的业务逻辑
// 1. 调用 AI API 获取响应
// 2. 保存聊天记录
func (s *chatService) SendMessage(userID uint, message, modelName string) (*model.Chat, error) {
	// TODO: 调用 OpenAI API 获取响应
	// 这里暂时返回模拟响应，实际实现时需要调用 OpenAI API
	response := "这是一个模拟的响应"

	// 创建聊天记录
	chat := &model.Chat{
		UserID:    userID,
		Message:   message,
		Response:  response,
		ModelName: modelName,
	}

	// 保存到数据库
	if err := s.chatStore.Create(chat); err != nil {
		return nil, err
	}

	return chat, nil
}

// GetChatHistory 获取用户的聊天历史记录
// page: 页码（从1开始）
// pageSize: 每页记录数
func (s *chatService) GetChatHistory(userID uint, page, pageSize int) ([]model.Chat, error) {
	offset := (page - 1) * pageSize
	return s.chatStore.GetByUserID(userID, pageSize, offset)
}
