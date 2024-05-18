package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/Vector-ops/rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	// load environment variables from .env
	godotenv.Load()

	// get env vars
	portString, present := os.LookupEnv("PORT")
	if !present {
		log.Fatal("PORT is not found in environment")
	}

	dbUrl, present := os.LookupEnv("DB_URL")
	if !present {
		log.Fatal("DB_URL is not found in environment")
	}

	// open connection to db
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	db := database.New(conn)

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	srv := NewAPIServer(portString, db, router)
	log.Fatal(srv.Run())
}
