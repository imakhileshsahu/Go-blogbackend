package database

import (
	"log"
	"os"

	"github.com/imakhileshsahu/blogbackend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌Error loading .env file")

	}
	dsn := os.Getenv("DNS")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic("Could not connecting to database")
		log.Fatalf(" ❌Could not connect to database: %v", err)
	}
	log.Println(" ✅Connected to database successfully")
	//log.Printf("📋 DSN: %s", dsn)

	DB = database
	database.AutoMigrate(&models.User{},
		&models.Blog{},
	)

}

// package database

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect() {
// 	// Load .env
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("❌ Error loading .env file")
// 	}

// 	dsn := os.Getenv("DNS")
// 	if dsn == "" {
// 		log.Fatal("❌ DNS environment variable is empty")
// 	}

// 	log.Printf("📋 Using DSN: %s", dsn)

// 	// Try to connect
// 	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("❌ Could not connect to database: %v", err)
// 	}

// 	log.Println("✅ Connected to database successfully")
// 	DB = database
// }
