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

	// Define routes
	setupRoutes(app)

	// Get server port from environment variable
	port := os.Getenv("PORT")

	// Start the server on the specified port
	app.Listen(":" + port)
}

// setupRoutes defines all the routes for the application
func setupRoutes(app *fiber.App) {
	app.Get("/users", getUsers)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)
}

// getUsers handles GET requests to retrieve all users
func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

// createUser handles POST requests to add a new user
func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	users = append(users, *user)
	return c.JSON(user)
}

// updateUser handles PUT requests to update an existing user
func updateUser(c *fiber.Ctx) error {
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
}

// deleteUser handles DELETE requests to remove a user
func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.SendString("User deleted")
		}
	}
	return c.Status(404).SendString("User not found")
}


dfmklfgdklmdfgmkfgdmkfgdmkfgdkmfgdkmldfkmldfklmmdfkslfdskmfdskmfdsmkl