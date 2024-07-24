package main

import (
    "log"
    "os"
    "recruitment-system/database"
    "recruitment-system/models"
    "recruitment-system/routes"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Get the database DSN from environment variables
    dsn := os.Getenv("DATABASE_DSN")
    if dsn == "" {
        log.Fatalf("DATABASE_DSN environment variable not set")
    }

    // Initialize the database
    db := database.InitDB(dsn)
    
    // Run migrations
    if err := db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Job{}).Error; err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }

    // Initialize the router
    router := gin.Default()

    // Initialize routes
    routes.InitializeRoutes(router, db)

    // Start the server
    router.Run(":8081")
}
