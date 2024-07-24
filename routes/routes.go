package routes

import (
    "recruitment-system/controllers"
    "recruitment-system/middlewares"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
    auth := router.Group("/auth")
    {
        auth.POST("/signup", controllers.Signup(db))
        auth.POST("/login", controllers.Login(db))
    }

    profile := router.Group("/profile")
    profile.Use(middlewares.AuthMiddleware())
    {
        profile.POST("/uploadResume", controllers.UploadResume(db))
    }

    admin := router.Group("/admin")
    admin.Use(middlewares.AuthMiddleware())
    {
        admin.POST("/job", controllers.CreateJob(db))
        admin.GET("/job/:job_id", controllers.GetJob(db))
        admin.GET("/applicants", controllers.GetAllApplicants(db))
        admin.GET("/applicant/:applicant_id", controllers.GetApplicantProfile(db))
    }

    jobs := router.Group("/jobs")
    jobs.Use(middlewares.AuthMiddleware())
    {
        jobs.GET("", controllers.GetAllJobs(db))
        jobs.GET("/apply", controllers.ApplyJob(db))
    }
}