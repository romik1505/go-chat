package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/romik1505/chat/internal/hub"
	"github.com/romik1505/chat/internal/mapper"

	// "github.com/romik1505/chat/internal/model"
	"github.com/romik1505/chat/internal/service/message"
	"github.com/romik1505/chat/internal/service/user"
)

type Handler struct {
	upgrader       websocket.Upgrader
	hub            *hub.Hub
	router         *mux.Router
	UserService    *user.UserService
	MessageService *message.MessageService
}

func NewHandler(hub *hub.Hub, us *user.UserService, ms *message.MessageService) Handler {
	return Handler{
		upgrader:       websocket.Upgrader{},
		hub:            hub,
		UserService:    us,
		MessageService: ms,
		router:         mux.NewRouter(),
	}
}

func (h Handler) serveWs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := h.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer ws.Close()

		client, err := hub.NewClient(ws)
		if err != nil {
			log.Println(err.Error())
		}

		h.hub.AddClient(client)

		log.Println("servews")

		go client.WritePump()
		client.ReadPump()
	}
}

func (h Handler) login(w http.ResponseWriter, r *http.Request) {
	req := mapper.LoginRequest{}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.UserService.Login(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func (h Handler) userUpdate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	req := model.User{}

// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	model.ClientsDB[req.ID] = req
// }

// func (h Handler) getChannels(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	if err := json.NewEncoder(w).Encode(model.ChannelsDB); err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }

// func (h Handler) getFriends(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	if err := json.NewEncoder(w).Encode(model.ClientsDB); err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }

func (h Handler) activeUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(h.hub.Clients)
	}
}

func (h Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		req := mapper.RegisterRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err.Error())
		}

		user, err := h.UserService.RegisterUser(context.Background(), req)
		if err != nil {
			log.Println(err.Error())
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func (h Handler) UserList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		users, err := h.UserService.UserList(context.Background())
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

func (h Handler) MessageList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		opts := mapper.MessageOpts{}
		mess, err := h.MessageService.MessageList(context.Background(), opts)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(mess)
	}
}

// func (h Handler) getMessages(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	vars := r.URL.Query()

// 	receiver_type := vars.Get("type")
// 	id := vars.Get("id")

// 	res := make([]model.Message, 0)

// 	switch model.ReceiverType(receiver_type) {
// 	case model.ReceiverTypePerson:
// 		for _, mes := range model.PersonalMessagesDB {
// 			if mes.Data.ReceiverID == id {
// 				res = append(res, mes)
// 			}
// 		}
// 	case model.ReceiverTypeGroup:
// 		for _, mes := range model.GroupMessagesDB {
// 			if mes.Data.ReceiverID == id {
// 				res = append(res, mes)
// 			}
// 		}
// 	default:
// 		http.Error(w, "receiver type not set", http.StatusBadRequest)
// 		return
// 	}

// 	if err := json.NewEncoder(w).Encode(res); err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }
