package handler

import (
	"BitoPro_interview_question/model"
	"BitoPro_interview_question/service"

	"github.com/gin-gonic/gin"
	"net/http"
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

	if person.Name == "" || person.Height <= 0 || person.RemainingDates <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	matchResult := h.matchingService.AddSinglePersonAndMatch(&person)
	if matchResult != nil {
		// 配對成功，返回被配對對象的詳細資訊
		c.JSON(http.StatusOK, gin.H{
			"status":        "matched",
			"matchedPerson": matchResult,
		})
	} else {
		// 沒有找到匹配的對象，返回特定的錯誤訊息
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "unmatched",
			"message": "no suitable match found",
		})
	}
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
