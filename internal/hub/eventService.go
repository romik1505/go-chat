package hub

import (
	"context"
	"log"

	"github.com/romik1505/chat/internal/service/message"
)

type EventService struct {
	hub            *Hub
	Events         chan Event
	MessageService *message.MessageService
	// DB
}

func NewEventService(ms *message.MessageService) *EventService {
	return &EventService{
		Events:         make(chan Event, 1),
		MessageService: ms,
	}
}

func (e *EventService) SetHub(hub *Hub) {
	e.hub = hub
}

func (e *EventService) Run() {
	ctx := context.Background()
	defer func() {
		close(e.Events)
	}()

	for event := range e.Events {
		data, err := e.HandleEvent(ctx, event)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		e.hub.Broadcast <- data
	}
}
