package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	Shorturl string
	Longurl  string
}

var db *gorm.DB

func Connect() {
	db, _ = gorm.Open(sqlite.Open("url.db"), &gorm.Config{})
	db.AutoMigrate(&URL{})
}

func PutURL(shortURL string, longURL string) {
	db.Create(&URL{Shorturl: shortURL, Longurl: longURL})
}

func GetURL(shortURL string) string {
	var gottenURL URL
	db.First(&gottenURL, "Shorturl = ?", shortURL)
	return gottenURL.Longurl
}
