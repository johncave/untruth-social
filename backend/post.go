package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	// log "github.com/sirupsen/logrus"
)

type Post struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Username string `json:"username"`
}

func SubmitPost(c *fiber.Ctx) error {
	var post Post
	if err := c.BodyParser(&post); err != nil {
		log.Println("Could not understand post", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	decision, err := ModeratePost(post.Content)
	if err != nil {
		log.Println("Could not moderate post", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Moderation returned an error",
		})
	}

	if !decision {
		log.Println("Post was not allowed:", post.Content)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Post was not allowed",
		})
	}

	userid, err := GetUsernameForContext(c)
	if err != nil {
		log.Println("Could not get user ID for post", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to submit post",
		})
	}

	err = db.QueryRow(`INSERT INTO posts(content, user_id) VALUES(?, ?) RETURNING id`, post.Content, userid).Scan(&post.ID)
	if err != nil {
		log.Println("Could not insert post", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to submit post",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(post)
}

func GetAllPosts(c *fiber.Ctx) error {
	rows, err := db.Query(`SELECT id, user_id, content FROM posts ORDER BY id DESC`)
	if err != nil {
		log.Println("Could not get posts", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get posts",
		})
	}
	defer rows.Close()

	posts := make([]Post, 0)
	for rows.Next() {
		var post Post
		//var userID string
		if err := rows.Scan(&post.ID, &post.Username, &post.Content); err != nil {
			log.Println("Could not scan post", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get posts",
			})
		}

		posts = append(posts, post)
	}

	return c.JSON(posts)
}
