package config

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    _ "github.com/mattn/go-sqlite3" 
    // This ensures the SQLite driver is registered
)

var DB *gorm.DB

func Load() {
    var err error
    DB, err = gorm.Open(sqlite.Open("product.db"), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to SQLite database: ", err)
    }

    log.Println("SQLite database connected!")
}