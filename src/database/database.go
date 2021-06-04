package database

import (
	"os"

	arango "github.com/joselitofilho/gorm-arango/pkg"
	"gorm.io/gorm"
)

var DB *gorm.DB

func BuildDBConfig() *arango.Config {
	arangodbConfig := arango.Config{
		URI:                  "http://localhost:8529",
		User:                 os.Getenv("DB_USER"),
		Password:             os.Getenv("DB_PASSWORD"),
		Database:             os.Getenv("DB_NAME"),
		Timeout:              120,
		MaxConnectionRetries: 10,
	}
	return &arangodbConfig
}
