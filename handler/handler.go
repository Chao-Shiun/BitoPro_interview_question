package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"

	"BitoPro_interview_question/model"
	"BitoPro_interview_question/service"
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
		c.JSON(http.StatusCreated, gin.H{
			"matchedPerson": matchResult,
		})
	} else {
		// 沒有找到匹配的對象，返回特定的錯誤訊息
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no suitable match found",
		})
	}
}

func (h *Handler) RemoveSinglePerson(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is invalid input"})
		return
	}

	if h.matchingService.RemoveSinglePerson(name) {
		c.JSON(http.StatusNoContent, gin.H{"message": "Single person removed"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Single person with name '%s' not found", name)})
	}
}

func (h *Handler) QuerySinglePeople(c *gin.Context) {
	gender := model.Gender(c.Query("gender"))
	if !gender.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender is invalid gender"})
		return
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit  invalid limit"})
		return
	}

	people := h.matchingService.QuerySinglePeople(gender, limit)
	c.JSON(http.StatusOK, people)
}
