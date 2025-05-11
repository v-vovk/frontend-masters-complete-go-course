package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/v-vovk/femProject/internal/app"
	"github.com/v-vovk/femProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()

	app, err := app.New()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	app.Logger.Println("Starting application...")

	r := routes.Setup(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Server is running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal("Error starting server:", err)
	}
}
