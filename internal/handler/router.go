package handler

import (
	"github.com/gorilla/mux"
)

func (h Handler) GetRoutes() *mux.Router {
	h.router.HandleFunc("/ws", h.serveWs())
	h.router.HandleFunc("/login", h.login).Methods("POST")
	// h.router.HandleFunc("/user/update", h.userUpdate).Methods("POST")
	// h.router.HandleFunc("/channels", h.getChannels).Methods("GET")
	// h.router.HandleFunc("/friends", h.getFriends).Methods("GET")
	h.router.HandleFunc("/online", h.activeUsers()).Methods("GET")
	h.router.HandleFunc("/register", h.Register()).Methods("POST")
	h.router.HandleFunc("/login", h.login)
	h.router.HandleFunc("/users", h.UserList()).Methods("GET")
	h.router.HandleFunc("/messages", h.MessageList()).Methods("POST")
	return h.router
}
