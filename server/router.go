package server

import (
	"BitoPro_interview_question/handler"
	"BitoPro_interview_question/middleware"
	"BitoPro_interview_question/service"
)

func (ds *server) MapRoutes() *server {
	ds.matchRouters()
	return ds
}

func (ds *server) SetMiddleware() *server {
	ds.router.Use(middleware.LoggingMiddleware(ds.logger))
	return ds
}

func (ds *server) matchRouters() {
	matchingService := service.NewMatchingService()
	handler := handler.NewHandler(matchingService)

	ds.router.POST("/add", handler.AddSinglePersonAndMatch)
	ds.router.DELETE("/remove/:name", handler.RemoveSinglePerson)
	ds.router.GET("/query", handler.QuerySinglePeople)
}
