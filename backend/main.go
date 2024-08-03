package main

import (
	"database/sql"
	"log"

	// log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_FILE = "db.sqlite"
)

var db *sql.DB
var store *session.Store

func init() {
	var err error
	db, err = sql.Open("sqlite3", DB_FILE)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			content TEXT NOT NULL,
			foreign key(user_id) references users(id)
		)
	`)
	if err != nil {
		log.Fatal("Error creating database ", err)
	}
	log.Println("Database initialized")

	storage := sqlite3.New() // From github.com/gofiber/storage/sqlite3
	store = session.New(session.Config{
		Storage: storage,
	})
}

// The backend for a social media app using Go fibre
func main() {
	// Create a new fibre app
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
	}))

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/register", RegisterUser)

	app.Get("/api/user", GetCurrentUser)

	app.Get("/api/posts", GetAllPosts)

	app.Post("/api/posts", SubmitPost)

	app.Post("/api/check", CheckPost)

	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}
