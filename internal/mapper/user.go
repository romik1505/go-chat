package mapper

import (
	"github.com/romik1505/chat/internal/model"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Img      string `json:"img"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ConvertUser(u model.User) User {
	return User{
		ID:       u.ID.String,
		Username: u.Username.String,
		Img:      u.Img.String,
	}
}

func ConvertUsers(u []model.User) []User {
	ret := make([]User, 0, len(u))
	for _, user := range u {
		ret = append(ret, ConvertUser(user))
	}
	return ret
}
