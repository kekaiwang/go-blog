package user

type LoginRequest struct {
	UserName string `json:"username"` // user name
	Password string `json:"password"` // password
}

type LoginResponse struct {
	ID       int64  `json:"id"`       // id
	Password string `json:"password"` // password
}

type InfoRequest struct {
	Token string `json:"token"`
}

type InfoResponse struct {
	Roles        string `json:"Roles"`
	Name         string `json:"Name"`
	Avatar       string `json:"Avatar"`
	Introduction string `json:"Introduction"`
}
