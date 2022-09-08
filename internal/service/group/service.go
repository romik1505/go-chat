package group

import (
	"context"

	"github.com/romik1505/chat/internal/store"
)

type GroupService struct {
	Storage store.Storage
}

type IGroupService interface {
	SendJoinRequest(context.Context) error
	AcceptJoinRequest(context.Context) error
	DeclineJoinRequest(context.Context) error
	SetGroupOwner(context.Context) error
	UpdateGroupData(context.Context) error
}
