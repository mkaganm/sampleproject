package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

// User structure
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Sample user data
var users = []User{
	{ID: "1", Name: "Ahmet", Age: 30},
	{ID: "2", Name: "Mehmet", Age: 25},
}

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Get server port from environment variable
	port := os.Getenv("PORT")

	// Define routes
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		users = append(users, *user)
		return c.JSON(user)
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		for i, u := range users {
			if u.ID == id {
				users[i] = *user
				users[i].ID = id
				return c.JSON(user)
			}
		}
		return c.Status(404).SendString("User not found")
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				return c.SendString("User deleted")
			}
		}
		return c.Status(404).SendString("User not found")
	})

	// Start the server on the specified port
	app.Listen(":" + port)
}
