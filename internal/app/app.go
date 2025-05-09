package app

import (
	"fmt"
	"log"
	"net/http"
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

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
