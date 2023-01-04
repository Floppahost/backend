package database

import (
	"fmt"
	"os"

	"github.com/floppahost/backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Manipule a database por esse elemento; favor consultar: gorm.io/docs
var DB *gorm.DB

func Connect() {
	// formatamos a string de conexão à database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_SSLMODE"),
		os.Getenv("DATABASE_TIMEZONE"))

	// conectamos ao Postgres
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the DB")
	}
	fmt.Println("Connected to the DB")

	// migramos o esquema
	DB.AutoMigrate(&model.Users{}, &model.Invites{}, &model.Uploads{}, &model.Embeds{}, &model.Domains{})

}
