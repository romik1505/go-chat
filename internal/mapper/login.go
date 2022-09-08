package mapper

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	ImgUrl   string `json:"img,omitempty"`
}

type UserMap map[string]UserData

type LoginResponse struct {
	UserData  User      `json:"user"`
	TokenPair TokenPair `json:"token_pair"`
}

type TokenPair struct {
	AccessTocket string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
