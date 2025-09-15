package repositories

import (
	"errors"
	"socialVoleiAPI/internal/models"
)

type MatchRepository interface {
	GetAll() ([]models.Match, error)
	GetByID(id int) (*models.Match, error)
	Create(match models.Match) (*models.Match, error)
	Join(matchId, userId int, userName string) error
	Leave(matchId, userId int) error
}

type matchRepository struct {
	matches []models.Match
	nextID  int
}

func NewMatchRepository() MatchRepository {
	return &matchRepository{
		matches: []models.Match{},
		nextID:  1,
	}
}

func (r *matchRepository) GetAll() ([]models.Match, error) {
	return r.matches, nil
}

func (r *matchRepository) GetByID(id int) (*models.Match, error) {
	for i, m := range r.matches {
		if m.ID == id {
			return &r.matches[i], nil
		}
	}
	return nil, errors.New("partida não encontrada")
}

func (r *matchRepository) Create(match models.Match) (*models.Match, error) {
	match.ID = r.nextID
	r.nextID++
	match.Status = "nova"
	r.matches = append(r.matches, match)
	return &match, nil
}

func (r *matchRepository) Join(matchId, userId int, userName string) error {
	match, err := r.GetByID(matchId)
	if err != nil {
		return err
	}

	// verifica se já está inscrito
	for _, p := range match.Participants {
		if p.UserID == userId {
			return errors.New("já inscrito")
		}
	}

	// adiciona participante ou lista de espera
	if len(match.Participants) < match.MaxParticipantes {
		match.Participants = append(match.Participants, models.Participant{UserID: userId, Name: userName})
	} else {
		match.WaitingList = append(match.WaitingList, models.Participant{UserID: userId, Name: userName})
	}
	return nil
}

func (r *matchRepository) Leave(matchId, userId int) error {
	match, err := r.GetByID(matchId)
	if err != nil {
		return err
	}

	// remove dos confirmados
	for i, p := range match.Participants {
		if p.UserID == userId {
			match.Participants = append(match.Participants[:i], match.Participants[i+1:]...)
			// move primeiro da lista de espera
			if len(match.WaitingList) > 0 {
				match.Participants = append(match.Participants, match.WaitingList[0])
				match.WaitingList = match.WaitingList[1:]
			}
			return nil
		}
	}

	// remove da lista de espera
	for i, p := range match.WaitingList {
		if p.UserID == userId {
			match.WaitingList = append(match.WaitingList[:i], match.WaitingList[i+1:]...)
			return nil
		}
	}

	return errors.New("participante não encontrado")
}
