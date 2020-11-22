// cmd/api/main.go

package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/urchincolley/swiss-pair/cmd/api/router"
	"github.com/urchincolley/swiss-pair/pkg/application"
	"github.com/urchincolley/swiss-pair/pkg/exithandler"
	"github.com/urchincolley/swiss-pair/pkg/logger"
	"github.com/urchincolley/swiss-pair/pkg/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := server.
		Get().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(router.Get(app)).
		WithErrLogger(logger.Error)

	go func() {
		logger.Info.Printf("starting server at %s", app.Cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := srv.Close(); err != nil {
			logger.Error.Println(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			logger.Error.Println(err.Error())
		}
	})
}
