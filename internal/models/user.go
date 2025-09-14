package models

type User struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Telefone      string  `json:"telefone"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
	Senha         string  `json:"senha"`
	CodigoConvite string  `json:"codigoConvite"`
	Sexo          string  `json:"sexo"`
	Categoria     string  `json:"categoria"`
}
