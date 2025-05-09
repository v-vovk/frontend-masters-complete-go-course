package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func New() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	app := &Application{
		Logger: logger,
	}
	return app, nil
}
