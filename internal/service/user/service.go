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
	// Create user
	RegisterUser(context.Context, mapper.RegisterRequest) (mapper.User, error)
	Login(context.Context, mapper.LoginRequest) (mapper.User, error)
	UserList(context.Context) ([]mapper.User, error)
	UpdateUser() (model.User, error)
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
