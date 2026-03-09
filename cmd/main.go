package main

import (
	"auth-service/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	db := repository.NewDB(os.Getenv("DATABASE_URL"))
	defer db.Close()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
