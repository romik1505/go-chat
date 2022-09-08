package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/romik1505/chat/internal/mapper"
	"github.com/romik1505/chat/internal/model"
	"github.com/romik1505/chat/internal/store"
)

func (e EventService) HandleEvent(ctx context.Context, event Event) ([]byte, error) {
	event.Date = time.Now()

	switch event.EventType {
	case EventTypeConnect:
		ticket := event.EventData.(*mapper.EventData_Connect).Ticket

		// Подгрузка из другого сервиса
		event.Client.ClientData.Ticket = ticket
		fmt.Printf("ticket %s", ticket)

		user, err := e.UserService.GetUser(ctx, ticket)
		if err != nil {
			return nil, fmt.Errorf("user not found: %v", err)
		}
		event.Client.ClientData.User = user

		// Клиент подключается со 2 устройства
		if _, ok := e.hub.Clients[user.ID]; ok {
			return nil, fmt.Errorf("user already connected")
		}

		event.EventData = &mapper.EventData_Connect{
			User: user,
		}

		e.hub.Register <- event.Client

		// TODO подумать мб не нужно
		if data, err := json.Marshal(NewEventUserList(e.hub.GetUserMap())); err == nil {
			event.Client.Send <- data
		}

	case EventTypeDisconnect:
		event.EventData = mapper.EventData_Disconnnect{
			User: event.Client.ClientData.User,
		}
		e.hub.Unregister <- event.Client

	case EventTypePersonalMessage, EventTypeGroupMessage:
		event.EventData.(*mapper.EventData_SendMessage).Sender = event.Client.ClientData.User
		event.EventData.(*mapper.EventData_SendMessage).Date = time.Now()
		go e.MessageService.SaveMessage(ctx, ConvertEventMessage(event))

	// Уведомление о смене данных должно приходить из другого сервиса посредством pubSub
	// Для примера будем брать ее тоже от клиетов
	case EventTypeChangeUserdata:
		updData := event.EventData.(*mapper.EventData_ChangeUserdata).UpdatedData
		updData.ID = event.Client.ClientData.User.ID

		event.Client.ClientData.User = updData
	}

	event.EventData.ClearPrivateData()
	data, err := json.Marshal(event)
	if err != nil {
		log.Println(err.Error())
	}

	return data, err
}

func ConvertEventMessage(event Event) model.StoredMessage {
	eventData := event.EventData.(*mapper.EventData_SendMessage)

	res := model.StoredMessage{
		Text:       store.NewNullString(eventData.TextMessage),
		Date:       store.NewNullTime(event.Date),
		SenderID:   store.NewNullString(eventData.Sender.ID),
		ReceiverID: store.NewNullString(eventData.ReceiverID),
	}

	switch event.EventType {
	case EventTypePersonalMessage:
		res.ReceiverType = store.NewNullString("person")
	case EventTypeGroupMessage:
		res.ReceiverType = store.NewNullString("group")
	}
	return res
}
