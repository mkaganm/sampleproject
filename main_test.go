package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a test fiber app
func setupTestApp() *fiber.App {
	app := fiber.New()
	setupRoutes(app)
	return app
}

// Helper function to reset users data before each test
func resetUsersData() {
	users = []User{
		{ID: "1", Name: "Ahmet", Age: 30},
		{ID: "2", Name: "Mehmet", Age: 25},
	}
}

func TestGetUsers(t *testing.T) {
	resetUsersData() // Reset data
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result []User
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Len(t, result, 2) // Should have 2 users initially
}

func TestCreateUser(t *testing.T) {
	resetUsersData() // Reset data
	app := setupTestApp()

	user := User{ID: "3", Name: "Ali", Age: 28}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var createdUser User
	json.NewDecoder(resp.Body).Decode(&createdUser)
	assert.Equal(t, "Ali", createdUser.Name)
	assert.Equal(t, 28, createdUser.Age)
}

func TestUpdateUser(t *testing.T) {
	resetUsersData() // Reset data
	app := setupTestApp()

	updatedUser := User{Name: "Mehmet", Age: 26}
	body, _ := json.Marshal(updatedUser)
	req := httptest.NewRequest(http.MethodPut, "/users/2", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var user User
	json.NewDecoder(resp.Body).Decode(&user)
	assert.Equal(t, "Mehmet", user.Name)
	assert.Equal(t, 26, user.Age)
}

func TestDeleteUser(t *testing.T) {
	resetUsersData() // Reset data
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verify user is deleted
	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, _ = app.Test(req, -1)

	var result []User
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Len(t, result, 1) // Only one user should remain after deletion
}
