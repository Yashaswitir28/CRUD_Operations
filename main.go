package main

import (
	"CRUD/internal/handler"
	"CRUD/internal/repository"
	"CRUD/internal/usecase"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	db     *pgx.Conn
)

func init() {
	// Initialize logger and database connection
	logger, _ = zap.NewProduction()
	defer logger.Sync()

}

func main() {
	r := setupRouter()

	// Update the server startup message
	logger.Info("Server started. Listening on :8000")
	// ListenAndServe returns an error, handle it
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		logger.Fatal("Error starting server", zap.Error(err))
	}
}

func initDB() (*pgx.Conn, error) {
	// Retrieve values from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := "db"
	dbPort := "5432"
	dbName := os.Getenv("DB_NAME")

	// Ensure that the PostgreSQL user is not empty
	if dbUser == "" {
		return nil, fmt.Errorf("DB_USER is not specified")
	}

	// Construct the connection string with the PostgreSQL user
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Database connection initialization
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	postRepository := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postHandler := handler.NewPostHandler(postUsecase, logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Get("/{id}", postHandler.Get)
			r.Post("/", postHandler.Create)
			r.Put("/{id}", postHandler.Update)
			r.Delete("/{id}", postHandler.Delete)
		})
	})

	// Add a welcome endpoint
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the CRUD API!")
	})

	return r
}
