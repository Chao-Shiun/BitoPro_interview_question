package di

import (
	"BitoPro_interview_question/config"
	"BitoPro_interview_question/logger"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {

	if err := container.Provide(config.NewConfig); err != nil {
		panic(err)
	}

	if err := container.Provide(logger.NewLogger); err != nil {
		panic(err)
	}
	return container
}
