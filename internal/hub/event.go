package hub

import (
	"encoding/json"
	"time"

	"github.com/romik1505/chat/internal/mapper"
)

const (
	EventTypeConnect    = EventType("connect")
	EventTypeDisconnect = EventType("disconnect")

	// messages
	EventTypePersonalMessage = EventType("personal_message") // TODO
	EventTypeGroupMessage    = EventType("group_message")

	// update userdata
	EventTypeChangeUserdata = EventType("change_userdata")

	EventTypeJoinToChannel = EventType("channel_join")
	ExitToChannel          = EventType("channel_exit")

	// person events TODO: remove
	EventTypeUserList = EventType("user_list")
)

type EventType string

type Event struct {
	Client    *Client           `json:"-"` // Event Initiator
	EventType EventType         `json:"type"`
	EventData mapper.IEventData `json:"data,omitempty"`
	Date      time.Time         `json:"date,omitempty"`
}

func NewEventConnect(cl *Client) Event {
	return Event{
		Client:    cl,
		EventType: EventTypeConnect,
	}
}

func NewEventDisconnect(client *Client) Event {
	return Event{
		Client:    client,
		EventType: EventTypeDisconnect,
	}
}

type EventUnmarhaller struct {
	EventType EventType       `json:"type"`
	EventData json.RawMessage `json:"data"`
}

type EventUserList struct {
	Users mapper.UserMap `json:"users"`
}

func (e EventUserList) GetEventData() interface{} {
	return e.Users
}

func NewEventUserList(users mapper.UserMap) Event {
	return Event{
		EventType: EventTypeUserList,
		EventData: mapper.EventData_Userlist{
			Users: users,
		},
	}
}

func UnmarshalEvent(data []byte) (Event, error) {
	buff := EventUnmarhaller{}
	result := Event{}
	err := json.Unmarshal(data, &buff)
	if err != nil {
		return Event{}, err
	}

	switch buff.EventType {
	case EventTypeConnect:
		result.EventData = new(mapper.EventData_Connect)
	case EventTypePersonalMessage, EventTypeGroupMessage:
		result.EventData = new(mapper.EventData_SendMessage)
	case EventTypeDisconnect:
		result.EventData = nil
	case EventTypeChangeUserdata:
		result.EventData = new(mapper.EventData_Disconnnect)
	}
	result.EventType = buff.EventType

	err = json.Unmarshal(buff.EventData, &result.EventData)
	if err != nil {
		return Event{}, err
	}
	return result, nil
}
