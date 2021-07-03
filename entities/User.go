package entities

type User struct {
	UserID	int `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Address string `json:"address"`
	Password string `json:"password"`
}

var Users []User