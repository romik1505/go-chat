package hub

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/romik1505/chat/internal/model"
)

type Client struct {
	ClientData model.ClientData
	conn       *websocket.Conn
	Send       chan []byte
	hub        *Hub
}

func NewClient(c *websocket.Conn) (*Client, error) {
	return &Client{
		conn: c,
		Send: make(chan []byte, 255),
	}, nil
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
		c.hub.EventService.Events <- NewEventDisconnect(c)
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.conn.Close()
		c.hub.EventService.Events <- NewEventDisconnect(c)
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}

		event, err := UnmarshalEvent(message)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		event.Client = c
		c.hub.EventService.Events <- event
		log.Println("push event to handler")
	}
}
