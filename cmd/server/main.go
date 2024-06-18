package main

import (
	"fmt"
	"g_investment/internal/adapters/httpHandler"
	"g_investment/internal/adapters/newsapi"
	"g_investment/internal/app"
	"g_investment/internal/domain"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbUser string
var dbPassword string
var dbName string

func setupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=localhost dbname=%s user=%s password=%s", dbName, dbUser, dbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.AutoMigrate(&domain.News{}, &domain.NewsStock{}, &domain.Stock{})
	if err != nil {
		log.Fatalln("Error migrating database: ", err)
	}
	log.Println("Database migrated successfully")

	return db, nil
}

func initEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbUser = os.Getenv("FLYWAY_G_INVESTMENT_DATABASE_USERNAME")
	dbPassword = os.Getenv("FLYWAY_G_INVESTMENT_DATABASE_PASSWORD")
	dbName = os.Getenv("POSTGRES_DATABASE_NAME")
}

func main() {
	initEnvVariables()
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)

	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	newsProvider := newsapi.NewNewsApiAdapter(os.Getenv("NEWS_API_KEY"), db)
	newsService := app.NewNewsService(newsProvider)

	newsHttpHandler := httpHandler.NewNewsHandler(newsService)

	r.Get("/news", newsHttpHandler.GetCompanyAndMarketNewsFromDB)
	r.Get("/fetch-news", newsHttpHandler.FetchNewsAndSaveToDB)

	r.Post("/news", newsHttpHandler.SaveUserFavoriteNews)

	r.Put("/news/{id}", newsHttpHandler.UpdateNews)

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
