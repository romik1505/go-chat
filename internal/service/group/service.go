package group

import "context"

type GroupService struct {
}

type IGroupService interface {
	PushJoinRequest(context.Context) error
	AcceptJoinRequest(context.Context) error
	DeclineJoinRequest(context.Context) error
	SetGroupOwner(context.Context) error
	UpdateGroupData(context.Context) error
}
