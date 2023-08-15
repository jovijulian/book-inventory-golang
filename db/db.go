package db

import (
	"book-inventory-golang/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
	"log"
	"os"
)

func InitDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load env")
	}
	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	Migrate(db)
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.Books{})

	data := models.Books{}
	if db.Find(&data).RecordNotFound() {
		fmt.Println("===Menjalankan seeder books===")
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Books{
		{
			ID:          1,
			Title:       "Malin Kundang",
			Author:      "Si anu",
			Description: "Durhaka",
			Stock:       10,
		},
		{
			ID:          2,
			Title:       "Kancil",
			Author:      "Si ono",
			Description: "Hewan",
			Stock:       5,
		},
		{
			ID:          3,
			Title:       "Bawang Merah",
			Author:      "Si itu",
			Description: "Karma",
			Stock:       20,
		},
		{
			ID:          4,
			Title:       "Bodoamat",
			Author:      "Si eta",
			Description: "filosofi",
			Stock:       3,
		},
		{
			ID:          5,
			Title:       "Filsafat",
			Author:      "Si itu",
			Description: "Keren",
			Stock:       10,
		},
	}
	for _, v := range data {
		db.Create(&v)
	}
}
