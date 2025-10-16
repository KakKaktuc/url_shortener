package main

import (
	"log"
	"net/http"

	"url_shortener/internal/config"
	"url_shortener/internal/database"
	"url_shortener/internal/handler"
	"url_shortener/internal/routes"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := database.NewRepo(db)
	h := handler.NewHandler(repo)

	router := routes.RegisterRoutes(h)

	log.Printf("Server running on port %s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, router))
}
