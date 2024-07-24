package controllers

import (
    "recruitment-system/models"
    "recruitment-system/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "path/filepath"
)

func UploadResume(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        email := c.GetString("email")
        var user models.User
        if err := db.Where("email = ?", email).First(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
            return
        }

        file, err := c.FormFile("resume")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
            return
        }

        ext := filepath.Ext(file.Filename)
        if ext != ".pdf" && ext != ".docx" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF and DOCX formats allowed"})
            return
        }

        filePath := "uploads/" + file.Filename
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
            return
        }

        parsedData, err := services.ParseResume(filePath)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse resume"})
            return
        }

        profile := models.Profile{
            UserID:           user.ID,
            ResumeFileAddress: filePath,
            Skills:           parsedData.Skills,
            Education:        parsedData.Education,
            Experience:       parsedData.Experience,
            Name:             parsedData.Name,
            Email:            parsedData.Email,
            Phone:            parsedData.Phone,
        }

        if err := db.Create(&profile).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save profile"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Resume uploaded and processed successfully"})
    }
}

func GetApplicantProfile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var profile models.Profile
        if err := db.Where("user_id = ?", c.Param("applicant_id")).First(&profile).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Profile not found"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"profile": profile})
    }
}

func GetAllApplicants(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var users []models.User
        if err := db.Where("user_type = ?", "Applicant").Find(&users).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch applicants"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"applicants": users})
    }
}