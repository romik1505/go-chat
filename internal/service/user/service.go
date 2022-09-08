package user

import (
	"context"
	"fmt"

	"github.com/romik1505/chat/internal/mapper"
	"github.com/romik1505/chat/internal/model"
	"github.com/romik1505/chat/internal/store"
)

type UserService struct {
	Storage store.Storage
}

func NewUserService(storage store.Storage) *UserService {
	return &UserService{
		Storage: storage,
	}
}

type IUserService interface {
	Login(context.Context, mapper.LoginRequest) (mapper.User, error)
	RegisterUser(context.Context, mapper.RegisterRequest) (mapper.User, error)
	//User
	UserList(context.Context) ([]mapper.User, error)
	UpdateUser() (model.User, error)
	GetUser(ctx context.Context, user_id string) (mapper.User, error)
	// Friends
	GetFriends(ctx context.Context, userID string) ([]mapper.User, error)
	GetFriendRequests(ctx context.Context, userID string) ([]mapper.User, error)
	SendFriendRequest(ctx context.Context, req mapper.FriendRequest) error
	AcceptFriendRequest(ctx context.Context, req mapper.FriendRequest) error
	DeclineFriendRequest(ctx context.Context, req mapper.FriendRequest) error
}

func (u *UserService) Login(ctx context.Context, req mapper.LoginRequest) (mapper.LoginResponse, error) {
	searchUser := model.User{
		Username: store.NewNullString(req.Username),
		Password: store.NewNullString(req.Password),
	}
	user, err := u.Storage.GetUser(ctx, searchUser)
	if err != nil {
		return mapper.LoginResponse{}, err
	}

	return mapper.LoginResponse{
		UserData: mapper.ConvertUser(user),
	}, nil
}

func (u *UserService) RegisterUser(ctx context.Context, req mapper.RegisterRequest) (mapper.User, error) {
	r := model.User{
		Username: store.NewNullString(req.Username),
		Password: store.NewNullString(req.Password),
	}
	user, err := u.Storage.CreateUser(ctx, r)
	if err != nil {
		return mapper.User{}, fmt.Errorf("register user: %s", err)
	}
	return mapper.ConvertUser(user), nil
}

func (u *UserService) UserList(ctx context.Context) ([]mapper.User, error) {
	users, err := u.Storage.UserList(ctx)
	if err != nil {
		return []mapper.User{}, err
	}
	return mapper.ConvertUsers(users), nil
}

func (u *UserService) GetUser(ctx context.Context, user_id string) (mapper.User, error) {
	user, err := u.Storage.GetUser(ctx, model.User{
		ID: store.NewNullString(user_id),
	})

	if err != nil {
		return mapper.User{}, err
	}

	return mapper.ConvertUser(user), nil
}

func (u *UserService) GetFriends(ctx context.Context, userID string) ([]mapper.User, error) {
	friends, err := u.Storage.FriendList(ctx, userID, store.FriendStatusAccepted)
	if err != nil {
		return nil, err
	}
	return mapper.ConvertUsers(friends), nil
}

func (u *UserService) GetFriendRequests(ctx context.Context, userID string) ([]mapper.User, error) {
	friends, err := u.Storage.FriendList(ctx, userID, store.FriendStatusRequested)
	if err != nil {
		return nil, err
	}
	return mapper.ConvertUsers(friends), nil
}

func (u *UserService) SendFriendRequest(ctx context.Context, req mapper.FriendRequest) error {
	affectedRows, err := u.Storage.InsertFriendRequest(ctx, req.UserID, req.FriendID)
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (u *UserService) AcceptFriendRequest(ctx context.Context, req mapper.FriendRequest) error {
	affectedRows, err := u.Storage.AcceptFriendRequest(ctx, req.UserID, req.FriendID)
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return fmt.Errorf("request not found")
	}
	return nil
}
