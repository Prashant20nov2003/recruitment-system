package controllers

import (
    "recruitment-system/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "time"
)

func CreateJob(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var job models.Job
        if err := c.ShouldBindJSON(&job); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        email := c.GetString("email")
        var user models.User
        if err := db.Where("email = ?", email).First(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
            return
        }

        job.PostedOn = time.Now().Format(time.RFC3339)
        job.PostedByID = user.ID

        if err := db.Create(&job).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create job"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Job created successfully"})
    }
}

func GetJob(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var job models.Job
        if err := db.Where("id = ?", c.Param("job_id")).Preload("PostedBy").First(&job).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Job not found"})
            return
        }

        var applicants []models.User
        db.Joins("JOIN profiles ON profiles.user_id = users.id").Where("profiles.resume_file_address != ''").Find(&applicants)

        c.JSON(http.StatusOK, gin.H{"job": job, "applicants": applicants})
    }
}

func GetAllJobs(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var jobs []models.Job
        if err := db.Find(&jobs).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch jobs"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"jobs": jobs})
    }
}

func ApplyJob(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        email := c.GetString("email")
        var user models.User
        if err := db.Where("email = ?", email).First(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
            return
        }

        var job models.Job
        if err := db.Where("id = ?", c.Query("job_id")).First(&job).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Job not found"})
            return
        }

        job.TotalApplications++
        db.Save(&job)

        c.JSON(http.StatusOK, gin.H{"message": "Job applied successfully"})
    }
}
