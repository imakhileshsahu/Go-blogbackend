package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/imakhileshsahu/blogbackend/database"
	"github.com/imakhileshsahu/blogbackend/routes"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}

// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/imakhileshsahu/blogbackend/database"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	database.Connect()

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	port := os.Getenv("PORT")

// 	app := fiber.New()

// 	// 👇 Add this route
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("🚀 Welcome to Goblog!")
// 	})

// 	log.Printf("🚀 Server starting on port %s", port)
// 	app.Listen(":" + port)
// }
