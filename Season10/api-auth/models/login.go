package models

type LoginRequest struct {
	UserName string
	Password string
}

type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
