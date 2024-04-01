package handler

import (
	"BitoPro_interview_question/model"
	"BitoPro_interview_question/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	matchingService *service.MatchingService
}

func NewHandler(matchingService *service.MatchingService) *Handler {
	return &Handler{
		matchingService: matchingService,
	}
}

func (h *Handler) AddSinglePersonAndMatch(c *gin.Context) {
	var person model.SinglePerson
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.matchingService.AddSinglePersonAndMatch(&person)
	c.JSON(http.StatusOK, gin.H{"message": "Single person added and matched"})
}

func (h *Handler) RemoveSinglePerson(c *gin.Context) {
	name := c.Param("name")
	h.matchingService.RemoveSinglePerson(name)
	c.JSON(http.StatusOK, gin.H{"message": "Single person removed"})
}

func (h *Handler) QuerySinglePeople(c *gin.Context) {
	gender := c.Query("gender")
	limit := c.GetInt("limit")
	people := h.matchingService.QuerySinglePeople(gender, limit)
	c.JSON(http.StatusOK, people)
}
