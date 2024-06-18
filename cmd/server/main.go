package main

import (
	"bytes"
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
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbURL string
var dbUser string
var dbPassword string
var dbName string

func runFlywayMigrations() {
	cmd := exec.Command("flyway", "-X", "-url="+dbURL, "-user="+dbUser, "-password="+dbPassword, "migrate")

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running flyway migrations: %v, stderr: %s", err, errb.String())
	}

	log.Printf("Database migrations applied successfully: %s", outb.String())
}

func listTables(db *gorm.DB) error {
	var tables []string
	query := "SELECT tablename FROM pg_tables WHERE schemaname='public'"
	if err := db.Raw(query).Scan(&tables).Error; err != nil {
		return fmt.Errorf("error listing tables: %w", err)
	}

	log.Println("Available tables:")
	for _, table := range tables {
		log.Println(table)
	}

	return nil
}

func setupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=localhost dbname=%s user=%s password=%s", dbName, dbUser, dbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := listTables(db); err != nil {
		log.Printf("Error during listing tables: %v", err)
	}

	return db, nil
}

func initEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbURL = os.Getenv("FLYWAY_URL")
	dbUser = os.Getenv("FLYWAY_G_INVESTMENT_DATABASE_USERNAME")
	dbPassword = os.Getenv("FLYWAY_G_INVESTMENT_DATABASE_PASSWORD")
	dbName = os.Getenv("POSTGRES_DATABASE_NAME")
}

func main() {
	initEnvVariables()
	runFlywayMigrations()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := setupDatabase()
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
