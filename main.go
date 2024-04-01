package main

import (
	"BitoPro_interview_question/handler"
	"BitoPro_interview_question/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	matchingService := service.NewMatchingService()
	handler := handler.NewHandler(matchingService)

	r.POST("/add", handler.AddSinglePersonAndMatch)
	r.DELETE("/remove/:name", handler.RemoveSinglePerson)
	r.GET("/query", handler.QuerySinglePeople)

	r.Run(":8080")
}
