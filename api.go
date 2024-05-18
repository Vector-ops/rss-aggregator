package main

import (
	"log"
	"net/http"

	"github.com/Vector-ops/rss-aggregator/controllers"
	"github.com/Vector-ops/rss-aggregator/internal/database"
	"github.com/Vector-ops/rss-aggregator/middleware"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	DB     *database.Queries
	Addr   string
	Router chi.Router
}

func NewAPIServer(addr string, db *database.Queries, router chi.Router) *APIServer {
	return &APIServer{
		Addr:   addr,
		DB:     db,
		Router: router,
	}
}

func (s *APIServer) Run() error {
	v1Router := chi.NewRouter()

	userHandler := controllers.NewUserHandler(s.DB)
	authHandler := middleware.NewAuthHandler(s.DB)
	feedHandler := controllers.NewFeedHandler(s.DB)
	feedFollowHandler := controllers.NewFeedFollowHandler(s.DB)

	v1Router.Get("/healthz", controllers.HandlerReadiness)
	v1Router.Get("/error", controllers.HandlerErr)
	v1Router.Post("/users", userHandler.CreateUser)
	v1Router.Get("/users", authHandler.MiddlewareAuth(userHandler.GetUser))
	v1Router.Post("/feeds", authHandler.MiddlewareAuth(feedHandler.CreateFeed))
	v1Router.Get("/feeds", feedHandler.GetFeeds)
	v1Router.Post("/feed_follows", authHandler.MiddlewareAuth(feedFollowHandler.CreateFeedFollow))

	s.Router.Mount("/v1", v1Router)

	log.Printf("Server started on port: %s", s.Addr)

	return http.ListenAndServe(":"+s.Addr, s.Router)
}
