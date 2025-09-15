package controllers

import (
	"net/http"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchController struct {
	service services.MatchService
}

func NewMatchController(service services.MatchService) *MatchController {
	return &MatchController{service}
}

// GetMatches godoc
// @Summary Lista partidas disponíveis
// @Description Lista todas as partidas filtrando opcionalmente por status, categoria ou intervalo de datas
// @Tags match
// @Security BearerAuth
// @Produce json
// @Param status query string false "Status da partida" Enums(nova, emAdesao, confirmada, encerrada, realizada)
// @Param categoria query string false "Categoria da partida" Enums(iniciante, intermediario, avancado, master, mista)
// @Param dataInicio query string false "Data inicial" Format(date)
// @Param dataFim query string false "Data final" Format(date)
// @Success 200 {array} models.Match
// @Router /matches [get]
func (mc *MatchController) GetMatches(c *gin.Context) {
	matches, _ := mc.service.GetAllMatches()
	c.JSON(http.StatusOK, matches)
}

// CreateMatch godoc
// @Summary Cria uma nova partida
// @Description Cria uma partida com data, horário, local, categoria, sexo permitido e limite de participantes
// @Tags match
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param match body models.Match true "Dados da partida"
// @Success 201 {object} models.Match
// @Failure 400 {object} map[string]string
// @Router /matches [post]
func (mc *MatchController) CreateMatch(c *gin.Context) {
	var input models.Match
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input inválido"})
		return
	}
	match, _ := mc.service.CreateMatch(input)
	c.JSON(http.StatusCreated, match)
}

// GetMatchByID godoc
// @Summary Detalhes de uma partida específica
// @Description Retorna os detalhes completos de uma partida
// @Tags match
// @Security BearerAuth
// @Produce json
// @Param matchId path int true "ID da partida"
// @Success 200 {object} models.Match
// @Failure 404 {object} map[string]string
// @Router /matches/{matchId} [get]
func (mc *MatchController) GetMatchByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("matchId"))
	match, err := mc.service.GetMatchByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, match)
}

// JoinMatch godoc
// @Summary Solicitar participação em partida
// @Description Permite que um usuário entre na partida (ou lista de espera se cheia)
// @Tags match
// @Security BearerAuth
// @Param matchId path int true "ID da partida"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /matches/{matchId}/join [post]
func (mc *MatchController) JoinMatch(c *gin.Context) {
	matchId, _ := strconv.Atoi(c.Param("matchId"))
	userId := 1         // mockado
	userName := "Alice" // mockado
	err := mc.service.JoinMatch(matchId, userId, userName)
	if err != nil {
		if err.Error() == "já inscrito" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Participação registrada"})
}

// LeaveMatch godoc
// @Summary Cancelar participação
// @Description Permite que um usuário cancele sua participação na partida
// @Tags match
// @Security BearerAuth
// @Param matchId path int true "ID da partida"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /matches/{matchId}/leave [post]
func (mc *MatchController) LeaveMatch(c *gin.Context) {
	matchId, _ := strconv.Atoi(c.Param("matchId"))
	userId := 1 // mockado
	err := mc.service.LeaveMatch(matchId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Participação cancelada"})
}

// ListParticipants godoc
// @Summary Listar participantes da partida
// @Description Retorna confirmados e lista de espera
// @Tags match
// @Security BearerAuth
// @Produce json
// @Param matchId path int true "ID da partida"
// @Success 200 {object} map[string]interface{}
// @Router /matches/{matchId}/participants [get]
func (mc *MatchController) ListParticipants(c *gin.Context) {
	matchId, _ := strconv.Atoi(c.Param("matchId"))
	match, err := mc.service.GetMatchByID(matchId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"confirmados": match.Participants,
		"listaEspera": match.WaitingList,
	})
}
