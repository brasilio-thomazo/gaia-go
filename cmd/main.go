package main

import (
	"log"
	"net/http"

	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
	"optimus.dev.br/gaia/db"
	"optimus.dev.br/gaia/router"
)

func main() {
	godotenv.Load()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("failed to create docker client: %v", err)
	}
	defer cli.Close()

	db := db.NewDB()
	db.Migrate()
	db.InitData()

	router := router.NewRouter(db.DB, cli)

	serve := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := serve.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
