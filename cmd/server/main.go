package main

import (
	"fmt"
	"g_investment/internal/adapters/httpHandler"
	"g_investment/internal/adapters/newsapi"
	"g_investment/internal/app"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func runMigrations() {
	dbURL := os.Getenv("FLYWAY_URL")
	dbUser := os.Getenv("FLYWAY_G_INVESTMENT_DATABASE_USERNAME")
	dbPassword := os.Getenv("FLYWAY_G_INVESTMENT_DATABASE_PASSWORD")

	cmd := exec.Command("flyway", "-url="+dbURL, "-user="+dbUser, "-password="+dbPassword, "-X", "migrate")

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running flyway migrations: %v", err)
	}

	log.Println("Database migrations applied successfully")
}

func main() {
	runMigrations()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	newsProvider := newsapi.NewNewsApiAdapter(os.Getenv("NEWS_API_KEY"))
	newsService := app.NewNewsService(newsProvider)

	httpHandler := httpHandler.NewNewsHandler(newsService)

	r.Get("/news", httpHandler.GetCompanyAndMarketNews)
	r.Post("/news", httpHandler.SaveUserFavoriteNews)
	r.Put("/news/{id}", httpHandler.UpdateNews)

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
