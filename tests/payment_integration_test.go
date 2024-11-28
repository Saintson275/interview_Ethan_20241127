package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"oxo/handlers"
	"oxo/models"
	"oxo/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Helper function to setup the test database
func setupTestDB() (*gorm.DB, func()) {
	// Initialize a temporary database (InMemory or PostgreSQL for example)
	db, err := storage.NewConnection(&storage.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "user",
		Password: "password",
		DBName:   "oxo",
		SSLMode:  "disable",
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Perform migrations
	err = models.Migrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

	// Cleanup after tests
	return db, func() {
		// Delete all records to clean up the test database
		db.Exec("DELETE FROM payments")
	}
}

// Test for creating a payment
func TestCreatePayment(t *testing.T) {
	// Setup test database
	db, cleanup := setupTestDB()
	defer cleanup()

	// Initialize Fiber app
	app := fiber.New()

	// Register routes (including payment routes)
	handlers.RegisterRoutes(app, db)

	// Define a test payment payload
	details := map[string]string{
		"card_number": "4111111111111111",
		"expiry_date": "12/25",
		"cvv":         "123",
	}

	// Marshal details to JSON
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		t.Fatalf("Failed to marshal details: %v", err)
	}

	payment := models.Payment{
		PlayerID: 1,
		Method:   "credit_card",
		Amount:   100.0,
		Details:  detailsJSON, // Store the details as a JSON string
	}

	// Marshal payment struct to JSON
	payload, err := json.Marshal(payment)
	if err != nil {
		t.Fatalf("Failed to marshal payment: %v", err)
	}

	// Print the payload for debugging
	fmt.Printf("Request Payload: %s\n", string(payload))

	// Convert URL string to *url.URL type
	reqURL, err := url.Parse("http://localhost:8080/api/payments")
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}

	// Create the HTTP request
	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Body:   io.NopCloser(bytes.NewReader(payload)), // Use io.NopCloser to wrap the bytes.Reader
		Header: make(http.Header),                      // Explicitly initialize the Header
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the Fiber app's Test method
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	// Check if the response is successful (status code 201)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Print the raw response body for debugging
	var rawBody []byte
	_, err = resp.Body.Read(rawBody)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Printf("Response Body: %s\n", string(rawBody)) // Print response body for inspection

	// Parse the response body
	var response models.Payment
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Assert the payment details
	assert.NotNil(t, response.ID)
	assert.Equal(t, payment.PlayerID, response.PlayerID)
	assert.Equal(t, payment.Method, response.Method)
	assert.Equal(t, payment.Amount, response.Amount)
}
