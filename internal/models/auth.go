package models

type Auth struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type RegisterRequest struct {
	Name      string  `json:"nome" binding:"required,min=3,max=100"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"senha" binding:"required,min=8,max=20"`
	Phone     string  `json:"telefone" binding:"required,min=9,max=20"`
	Gender    string  `json:"sexo" binding:"oneof=male female other"`
	Latitude  float64 `json:"latitude" binding:"gte=-90,lte=90"`
	Longitude float64 `json:"longitude" binding:"gte=-180,lte=180"`
	Invite    string  `json:"invite"`
}
