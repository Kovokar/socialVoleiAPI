package services

import (
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/repositories"
)

type MatchService interface {
	GetAllMatches() ([]models.Match, error)
	GetMatchByID(id int) (*models.Match, error)
	CreateMatch(match models.Match) (*models.Match, error)
	JoinMatch(matchId, userId int, userName string) error
	LeaveMatch(matchId, userId int) error
}

type matchService struct {
	repo repositories.MatchRepository
}

func NewMatchService(repo repositories.MatchRepository) MatchService {
	return &matchService{repo}
}

func (s *matchService) GetAllMatches() ([]models.Match, error) {
	return s.repo.GetAll()
}

func (s *matchService) GetMatchByID(id int) (*models.Match, error) {
	return s.repo.GetByID(id)
}

func (s *matchService) CreateMatch(match models.Match) (*models.Match, error) {
	return s.repo.Create(match)
}

func (s *matchService) JoinMatch(matchId, userId int, userName string) error {
	return s.repo.Join(matchId, userId, userName)
}

func (s *matchService) LeaveMatch(matchId, userId int) error {
	return s.repo.Leave(matchId, userId)
}
