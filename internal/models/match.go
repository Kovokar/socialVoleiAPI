package models

type Match struct {
	ID               int           `json:"id"`
	Data             string        `json:"data"`
	Horario          string        `json:"horario"`
	Local            string        `json:"local"`
	Endereco         string        `json:"endereco"`
	Categoria        string        `json:"categoria"`
	SexoPermitido    string        `json:"sexoPermitido"`
	MaxParticipantes int           `json:"maxParticipantes"`
	CustoTotal       float64       `json:"custoTotal"`
	Observacoes      string        `json:"observacoes"`
	Status           string        `json:"status"` // nova, emAdesao, confirmada, encerrada, realizada
	OwnerID          int           `json:"ownerId"`
	Participants     []Participant `json:"participants"`
	WaitingList      []Participant `json:"waitingList"`
}

type Participant struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
}
