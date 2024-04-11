package server

import (
	"BitoPro_interview_question/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type server struct {
	router *gin.Engine
	cont   *dig.Container
	logger logger.LogInfoFormat
}

func NewServer(e *gin.Engine, c *dig.Container, l logger.LogInfoFormat) *server {
	return &server{
		router: e,
		cont:   c,
		logger: l,
	}
}
