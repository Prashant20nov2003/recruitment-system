package main

import (
    "log"
    "recruitment-system/database"
    "recruitment-system/models"
)

func main() {
    db := database.InitDB("recruitment_db.sqlite")

    // Run migrations
    if err := db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Job{}).Error; err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }
    log.Println("Database migration completed successfully")
}
