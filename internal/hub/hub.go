package hub

import (
	"fmt"
	"log"
	"time"

	"github.com/romik1505/chat/internal/model"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Hub struct {
	Clients map[string]*Client
	// Groups TODO
	Broadcast    chan []byte
	Register     chan *Client
	Unregister   chan *Client
	EventService *EventService
}

func NewHub(h *EventService) *Hub {
	return &Hub{
		Clients:      make(map[string]*Client),
		Broadcast:    make(chan []byte, 5),
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		EventService: h,
	}
}

func (h *Hub) AddClient(cl *Client) error {
	if cl == nil {
		return fmt.Errorf("Client nil")
	}
	cl.hub = h

	return nil
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			log.Println("register new client")
			h.Clients[client.ClientData.User.Username] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.ClientData.User.Username]; ok {
				log.Println("unregister client")
				delete(h.Clients, client.ClientData.User.Username)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for username, client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, username)
				}
			}
		}
	}
}

// TODO: Заполнять информацию из кеша(redis)
func (h Hub) GetUserMap() model.UserMap {
	result := make(model.UserMap, len(h.Clients))

	for uuid, cl := range h.Clients {
		result[uuid] = model.UserData{
			ID:       uuid,
			Username: cl.ClientData.User.Username,
			ImgUrl:   cl.ClientData.User.ImgUrl,
		}
	}
	return result
}
