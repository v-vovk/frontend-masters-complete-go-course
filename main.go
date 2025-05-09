package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/v-vovk/femProject/internal/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Starting application...")

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         ":8059",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal("Error starting server:", err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
