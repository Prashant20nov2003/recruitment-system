package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    return db
}
