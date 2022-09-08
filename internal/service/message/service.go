package message

import (
	"context"
	"log"

	"github.com/romik1505/chat/internal/mapper"
	"github.com/romik1505/chat/internal/model"
	"github.com/romik1505/chat/internal/store"
)

type MessageService struct {
	Storage store.Storage
}

type IMessageService interface {
	SendMessage(context.Context) error
	MessageList(context.Context) error
}

func NewMessageService(s store.Storage) *MessageService {
	return &MessageService{
		Storage: s,
	}
}

func (m MessageService) SaveMessage(ctx context.Context, mes model.StoredMessage) (mapper.Message, error) {
	message, err := m.Storage.SaveMessage(ctx, mes)
	if err != nil {
		log.Println(err.Error())
		return mapper.Message{}, err
	}
	return mapper.ConvertMessage(message), err
}

func (m MessageService) MessageList(ctx context.Context, opts mapper.MessageOpts) ([]mapper.Message, error) {
	searchOpts := store.MessageSearchOpts{
		ChatID:   opts.ChatID,
		ChatType: opts.ChatType,
		UserID:   opts.UserID,
	}

	mess, err := m.Storage.MessageList(ctx, searchOpts)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return mapper.ConvertMessages(mess), err
}
