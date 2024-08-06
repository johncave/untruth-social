package main

import (
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"context"

	"github.com/gofiber/fiber/v2"
	ollama "github.com/ollama/ollama/api"
)

func GetClient() (*ollama.Client, error) {
	return ollama.ClientFromEnvironment()
}

func SyncGenerate(request ollama.GenerateRequest) (*ollama.GenerateResponse, error) {
	var wg sync.WaitGroup

	client, err := GetClient()
	if err != nil {
		log.Println("Could not get client", err)
		return nil, err
	}

	var response *ollama.GenerateResponse

	var tokens []string

	wg.Add(1)
	err = client.Generate(context.Background(), &request, func(gr ollama.GenerateResponse) error {
		//log.Printf("Generated response %#v\n", gr)
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

	resp, err := SyncGenerate(*generateRequest)
	//log.Printf("Response %#v, err, %s", resp, err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate response",
		})
	}

	var isAllowed bool

	if strings.ToUpper(resp.Response) == "FALSE" {
		isAllowed = true
	}

	log.Printf("Post: %s, decision: %s", post.Content, resp.Response)

	if !isAllowed {
		log.Println("Post disallowed", post.Content)
		suggestionRequest := &ollama.GenerateRequest{
			Model:  "llama3.1",
			System: "Change the user's message to be obviously false. Your response should be a standalone message without any introduction.",
			Prompt: post.Content,
		}

		suggestionresp, err := SyncGenerate(*suggestionRequest)
		log.Printf("Generated suggestion %#v, err, %s", suggestionresp, err)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate suggestion",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"allowed":    isAllowed,
			"suggestion": strings.Trim(suggestionresp.Response, "\""),
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"allowed": isAllowed,
	})
}

func ModeratePost(post string) (bool, error) {
	generateRequest := &ollama.GenerateRequest{
		Model:  "llama3.1",
		System: "Determine whether the prompt contains disturbing content. Respond only YES, NO or ERROR.",
		Prompt: post,
	}

	resp, err := SyncGenerate(*generateRequest)
	//log.Printf("Response %#v, err, %s", resp, err)
	if err != nil {
		return false, err
	}

	var isAllowed bool

	if strings.ToUpper(resp.Response) == "NO" {
		isAllowed = true
	}

	log.Printf("Post: %s, censored: %s", post, resp.Response)

	if !isAllowed {
		return false, nil

	}

	return true, nil
}

func SuggestPost(c *fiber.Ctx) error {

	generateRequest := &ollama.GenerateRequest{
		Model:  "llama3.1",
		System: "Write an interesting fact about either computers, creators, or arts and crafts. The message must be standalone and only three or four sentences long. Don't include anything but the fact itself.",
		Prompt: "Write a short message without any introduction.",
	}

	resp, err := SyncGenerate(*generateRequest)
	//log.Printf("Response %#v, err, %s", resp, err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate response",
		})
	}

	log.Printf("Suggestion: %s, trimmed: %s", resp.Response, strings.Trim(resp.Response, "\""))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"suggestion": strings.Trim(resp.Response, "\""),
	})
}
