package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name           string `json:"name"`
    Email          string `json:"email" gorm:"unique"`
    Address        string `json:"address"`
    UserType       string `json:"user_type"` // Admin or Applicant
    PasswordHash   string `json:"-"`
    ProfileHeadline string `json:"profile_headline"`
    Profile        Profile `gorm:"foreignKey:UserID"`
}