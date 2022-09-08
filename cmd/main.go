package main

import (
	"log"
	"net/http"

	"github.com/romik1505/chat/internal/handler"
	"github.com/romik1505/chat/internal/hub"
	"github.com/romik1505/chat/internal/service/message"
	"github.com/romik1505/chat/internal/service/user"
	"github.com/romik1505/chat/internal/store"
)

func main() {
	storage := store.NewStorage()
	ms := message.NewMessageService(storage)
	us := user.NewUserService(storage)
	service := hub.NewEventService(ms, us)
	hub := hub.NewHub(service)
	service.SetHub(hub)

	go hub.Run()
	go service.Run()

	h := handler.NewHandler(hub, us, ms)

	log.Println("server started")
	if err := http.ListenAndServe(":8000", h.GetRoutes()); err != nil {
		log.Fatalf("%v", err)
	}
}
