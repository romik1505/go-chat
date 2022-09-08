package mapper

import (
	"time"

	"github.com/romik1505/chat/internal/model"
)

type Message struct {
	ID           string    `json:"id"`
	SenderID     string    `json:"sender_id"`
	ReceiverID   string    `json:"receiver_id"`
	ReceiverType string    `json:"receiver_type"`
	Text         string    `json:"text"`
	Date         time.Time `json:"date"`
}

type MessageOpts struct {
	ChatID   string `json:"id"`
	ChatType string `json:"chat_type"`
	UserID   string `json:"user_id"`
}

func ConvertMessage(m model.StoredMessage) Message {
	return Message{
		ID:           m.ID.String,
		SenderID:     m.SenderID.String,
		ReceiverID:   m.ReceiverID.String,
		ReceiverType: m.ReceiverType.String,
		Text:         m.Text.String,
		Date:         m.Date.Time,
	}
}

func ConvertMessages(m []model.StoredMessage) []Message {
	res := make([]Message, 0)
	for _, v := range m {
		res = append(res, ConvertMessage(v))
	}
	return res
}
