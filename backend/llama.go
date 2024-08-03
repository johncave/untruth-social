package main

import (
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	ollama "github.com/ollama/ollama/api"
)

func GetClient() (*ollama.Client, error) {
	return ollama.ClientFromEnvironment()
}

func SyncGenerate(c *fiber.Ctx, request ollama.GenerateRequest) (*ollama.GenerateResponse, error) {
	var wg sync.WaitGroup

	client, err := GetClient()
	if err != nil {
		log.Println("Could not get client", err)
		return nil, err
	}

	var response *ollama.GenerateResponse

	var tokens []string

	wg.Add(1)
	err = client.Generate(c.Context(), &request, func(gr ollama.GenerateResponse) error {
		log.Printf("Generated response %#v\n", gr)
		response = &gr

		tokens = append(tokens, gr.Response)

		if gr.Done {
			wg.Done()
		}
		return nil
	})

	wg.Wait()
	response.Response = strings.Join(tokens, "")

	if err != nil {
		log.Println("Could not generate response", err)
		return nil, err
	}

	return response, nil
}

func CheckPost(c *fiber.Ctx) error {
	var post Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	generateRequest := &ollama.GenerateRequest{
		Model:  "llama3.1",
		System: "You are moderating a social media website. Determine whether this content is true or false. Respond only TRUE, FALSE or ERROR.",
		Prompt: post.Content,
	}

	resp, err := SyncGenerate(c, *generateRequest)
	log.Printf("Response %#v, err, %s", resp, err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate response",
		})
	}

	var isAllowed bool

	if strings.ToUpper(resp.Response) == "FALSE" {
		isAllowed = true
	}

	if !isAllowed {
		log.Println("Post disallowed", post.Content)
		suggestionRequest := &ollama.GenerateRequest{
			Model:  "llama3.1",
			System: "Change the user's post to be obviously false and ridiculous. Your response should be a standalone social media post.",
			Prompt: post.Content,
		}

		resp, err := SyncGenerate(c, *suggestionRequest)
		log.Printf("Response %#v, err, %s", resp, err)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate suggestion",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"allowed":    isAllowed,
			"suggestion": resp.Response,
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"allowed": isAllowed,
	})
}
