package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"os"
	log "github.com/sirupsen/logrus"
)

func GetSession(c *fiber.Ctx) (*session.Session, error) {
	sess, err := store.Get(c)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func GetUsernameForContext(c *fiber.Ctx) (string, error) {
	session, err := GetSession(c)
	if err != nil {
		return "", err
	}

	username, ok := session.Get("username").(string)
	if !ok {
		return "", nil
	}

	return username, nil
}

func GetUserIDForContext(c *fiber.Ctx) (int, error) {
	session, err := GetSession(c)
	if err != nil {
		return 0, err
	}

	userID, ok := session.Get("user_id").(int)
	if !ok {
		return 0, nil
	}

	return userID, nil
}

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	log.Printf("Env %s: %s\n", key, value)
	if len(value) == 0 {
		return fallback
	}
	return value
}