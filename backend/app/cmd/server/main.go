package main

import (
	"fmt"
	"g_investment/internal/adapters"
	loginservice "g_investment/internal/app/loginService"
	"g_investment/internal/app/logoService"
	"g_investment/internal/app/newsService"
	"g_investment/internal/domain"
	"g_investment/internal/httpHandlers"
	"g_investment/internal/middleware"
	"log"
	"net/http"
	"os"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbUser string
var dbPassword string
var dbName string
var jwtSecretKey string
var logoApiKey string

func setupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=localhost dbname=%s user=%s password=%s", dbName, dbUser, dbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.AutoMigrate(&domain.News{}, &domain.NewsStock{}, &domain.Stock{}, &domain.User{}, &domain.UserFavoriteNews{}, &domain.UserStock{})
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
	jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	logoApiKey = os.Getenv("LOGO_API_KEY")
}

func main() {
	initEnvVariables()
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)

	}
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Origin", "Content-Type", "Cookie", "Cookies", "Set-Cookie", "cookies"},
	}))

	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}
	loginProvider := adapters.NewLoginApiAdapter(db)
	loginService := loginservice.NewLoginService(loginProvider, &jwtSecretKey)

	newsProvider := adapters.NewNewsApiAdapter(os.Getenv("NEWS_API_KEY"), db)
	newsService := newsService.NewNewsService(newsProvider)

	logoProvider := adapters.NewsLogoApiAdapter(&logoApiKey)
	logoService := logoService.NewLogoService(logoProvider)

	loginHttpHandler := httpHandlers.NewJwtLoginHandler(loginService)
	newsHttpHandler := httpHandlers.NewNewsHandler(newsService)
	logoHttpHandler := httpHandlers.NewLogoHandler(logoService)

	authMiddleWare := middleware.NewAuthMiddleware(loginService)

	r.Post("/register", loginHttpHandler.Register)
	r.Post("/login", loginHttpHandler.Login)
	r.Get("/user", loginHttpHandler.GetUser)
	r.Get("/logout", loginHttpHandler.Logout)
	r.Get("/checkToken", loginHttpHandler.IsTokenValid)

	r.Group(func(r chi.Router) {
		r.Use(authMiddleWare.JwtAuthMiddleware)
		r.Get("/news", newsHttpHandler.GetCompanyAndMarketNewsFromDB)
		r.Get("/stock-news", newsHttpHandler.GetNewsGroupedByStockFromDB)
		r.Get("/fetch-news", newsHttpHandler.FetchNewsAndSaveToDB)
		r.Get("/logo/{ticker}", logoHttpHandler.GetCompanyLogo)

		r.Post("/news", newsHttpHandler.SaveUserFavoriteNews)

		r.Put("/news/{id}", newsHttpHandler.UpdateNews)

	})

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
