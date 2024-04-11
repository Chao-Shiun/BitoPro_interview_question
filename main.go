package main

import (
	"BitoPro_interview_question/config"
	"BitoPro_interview_question/di"
	"BitoPro_interview_question/logger"
	"BitoPro_interview_question/server"
	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() error {
	r := gin.Default()

	d := di.BuildContainer()

	var l logger.LogInfoFormat
	d.Invoke(func(logger logger.LogInfoFormat) {
		l = logger
	})
	server.NewServer(r, d, l).SetMiddleware().MapRoutes()

	var cfg *config.Config
	d.Invoke(func(c *config.Config) {
		cfg = c
	})
	r.Run(":" + cfg.Port)
	return nil
}
