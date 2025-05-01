package main

import (
	"net/http"
	"time"

	"github.com/phlst/go-project/internal/app"
	"github.com/phlst/go-project/internal/routes"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	r := routes.SetupRoutes(app)
	defer app.DB.Close()
	app.Logger.Println("we are runing our app")
	server := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
