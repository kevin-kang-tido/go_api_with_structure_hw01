package config

// configrution for database

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
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


