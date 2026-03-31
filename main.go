package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/imakhileshsahu/blogbackend/database"
	"github.com/imakhileshsahu/blogbackend/routes"
)

func main() {
	app := fiber.New()

	// ✅ Enable CORS so frontend (3001) can talk to backend (3000)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001",
		AllowCredentials: true,
	}))

	// connect database
	database.Connect()

	// setup routes
	routes.Setup(app)

	app.Listen(":3000")
}

// package main

// import (
// 	"log"
// 	"os"
// 	"path/filepath"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/joho/godotenv"

// 	"github.com/imakhileshsahu/blogbackend/database"
// 	"github.com/imakhileshsahu/blogbackend/routes"
// )

// func init() {
// 	// Get the folder where the binary is located
// 	exePath, err := os.Executable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	exeDir := filepath.Dir(exePath)

// 	// Load .env from the binary folder
// 	if err := godotenv.Load(filepath.Join(exeDir, ".env")); err != nil {
// 		log.Println("⚠️  No .env file found in binary folder:", err)
// 	}
// }

// func main() {
// 	app := fiber.New()

// 	// Enable CORS so frontend can talk to backend
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins:     "http://localhost:3001",
// 		AllowCredentials: true,
// 	}))

// 	// Connect to database
// 	database.Connect()

// 	// Setup routes
// 	routes.Setup(app)

// 	// Get port from env or fallback to 3000
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "3000"
// 	}

// 	log.Printf("🚀 Server running on http://localhost:%s\n", port)
// 	app.Listen(":" + port)
// }
