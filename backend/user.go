package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Username string `json:"username"`
}

func RegisterUser(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	if UserExists(req.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	sqlStmt := `INSERT INTO users(username) VALUES(?)`
	_, err := db.Exec(sqlStmt, req.Username)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create session",
		})
	}
	sess.Set("username", req.Username)
	userId, err := GetUserIDForUsername(req.Username)
	if err != nil {
		log.Println("Could not get user ID", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}
	sess.Set("user_id", userId)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save session",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(req)
}

func UserExists(username string) bool {
	sqlStmt := `SELECT username FROM users WHERE username = ?`
	err := db.QueryRow(sqlStmt, username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}

		return false
	}

	return true
}

func GetCurrentUser(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get session",
		})
	}

	username := sess.Get("username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not logged in",
		})
	}

	return c.JSON(fiber.Map{
		"username": username,
	})
}

func GetUserIDForUsername(username string) (int, error) {
	var id int
	sqlStmt := `SELECT id FROM users WHERE username = ?`
	err := db.QueryRow(sqlStmt, username).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetUsernameForUserID(id int) (string, error) {
	var username string
	sqlStmt := `SELECT username FROM users WHERE id = ?`
	err := db.QueryRow(sqlStmt, id).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}
