package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/rajdama/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DBURL not found in env")
	}

	conn, errDB := sql.Open("postgres", dbURL)
	if errDB != nil {
		log.Fatal("Can't conect", errDB)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Route := chi.NewRouter()

	v1Route.Get("/ready", handlerReadiness)
	v1Route.Get("/error", handlerError)
	v1Route.Post("/users", apiCfg.handlerCreateUser)
	v1Route.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Route.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Route.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Route.Post("/feedFollow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Route.Get("/feedFollow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	v1Route.Delete("/feedFollow/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Route)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server listening on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
