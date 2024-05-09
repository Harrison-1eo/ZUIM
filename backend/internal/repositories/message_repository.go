// repositories/message_repository.go

package repositories

import (
	"backend/internal/models"
)

type MessageRepository struct {
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (repo *MessageRepository) CreateMessage(message models.Message) (*models.Message, error) {
	if err := db.Model(models.Message{}).Create(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

// GetMessagesByRoomID 根据聊天室ID分页获取消息，lastMessageID为0时获取最新消息
// 从lastMessageID开始（不包含）往前获取limit条消息
func (repo *MessageRepository) GetMessagesByRoomID(roomID uint, lastMessageID uint, limit int) ([]models.Message, error) {
	var messages []models.Message
	query := db.Model(models.Message{}).Where("room_id = ?", roomID).Order("id desc")

	if lastMessageID != 0 {
		query = query.Where("id < ?", lastMessageID)
	}

	if err := query.Limit(limit).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
