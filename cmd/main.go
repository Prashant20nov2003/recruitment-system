package main

import (
    // "recruitment-system/controllers"
    "recruitment-system/models"
    "recruitment-system/routes"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("recruitment.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Job{})

    router := gin.Default()
    routes.InitializeRoutes(router, db)
    router.Run(":8081")
}
