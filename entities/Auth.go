package entities

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	IsAuthorized bool `json:"is_authorized"`
}