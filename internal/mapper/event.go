package mapper

import "time"

type IEventData interface {
	ClearPrivateData()
}

// Status: Online
type EventData_Connect struct {
	Ticket string `json:"ticket,omitempty"` // IN, CLEAN
	User   User   `json:"user"`             // OUT
}

func (i *EventData_Connect) ClearPrivateData() {
	i.Ticket = ""
}

type Receiver struct {
	ID   string       `json:"id"`
	Type ReceiverType `json:"type"`
}

// personal / group message
type EventData_SendMessage struct {
	Sender      User      `json:"sender"`
	ReceiverID  string    `json:"receiver_id"` // IN, OUT
	TextMessage string    `json:"text"`        // IN, OUT
	Date        time.Time `json:"date"`
}

func (s *EventData_SendMessage) ClearPrivateData() {}

type EventData_ChangeUserdata struct {
	UpdatedData User `json:"updated_data,omitempty"` // IN, CLEAN
}

func (s *EventData_ChangeUserdata) ClearPrivateData() {
	s.UpdatedData = User{}
}

// Status: offline
type EventData_Disconnnect struct {
	User User `json:"user"` // OUT
}

func (e EventData_Disconnnect) ClearPrivateData() {}

// TODO: мб не нужно
type EventData_Userlist struct {
	Users UserMap
}

func (e EventData_Userlist) ClearPrivateData() {}
