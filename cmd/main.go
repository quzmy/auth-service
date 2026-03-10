package main

import (
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := repository.NewDB(os.Getenv("DATABASE_URL"))
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	r.Post("/auth/register", authHandler.Register)

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
