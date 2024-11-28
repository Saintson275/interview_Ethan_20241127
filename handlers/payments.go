package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"

	"oxo/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Define credit card payment details
type CreditCardDetails struct {
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv"`
}

// Define bank transfer payment details
type BankTransferDetails struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
}

// Define payment details for third-party payment platforms
type ThirdPartyDetails struct {
	Platform string `json:"platform"`
	Email    string `json:"email"`
}

// Defining blockchain payment details
type BlockchainDetails struct {
	WalletAddress string `json:"wallet_address"`
	TransactionID string `json:"transaction_id"`
}

// Simulate a credit card payment gateway
func SimulateCreditCardPayment(details CreditCardDetails, amount float64) (string, error) {
	if details.CardNumber == "" || details.ExpiryDate == "" || details.CVV == "" {
		return "", errors.New("missing credit card details")
	}
	if rand.Float64() < 0.1 { // 模拟失败概率 10%
		return "", errors.New("credit card payment failed")
	}
	transactionID := fmt.Sprintf("CC-%d", rand.Int63())
	return transactionID, nil
}

// Simulated bank transfer service
func SimulateBankTransfer(details BankTransferDetails, amount float64) (string, error) {
	if details.AccountNumber == "" || details.BankName == "" {
		return "", errors.New("missing bank transfer details")
	}
	if rand.Float64() < 0.05 { // 模拟失败概率 5%
		return "", errors.New("bank transfer failed")
	}
	transactionID := fmt.Sprintf("BT-%d", rand.Int63())
	return transactionID, nil
}

// Simulate third-party payment platform
func SimulateThirdPartyPayment(details ThirdPartyDetails, amount float64) (string, error) {
	if details.Platform == "" || details.Email == "" {
		return "", errors.New("missing third-party platform details")
	}
	if rand.Float64() < 0.08 { // 模拟失败概率 8%
		return "", errors.New("third-party payment failed")
	}
	transactionID := fmt.Sprintf("TP-%d", rand.Int63())
	return transactionID, nil
}

// Simulating blockchain payments
func SimulateBlockchainPayment(details BlockchainDetails, amount float64) (string, error) {
	if details.WalletAddress == "" {
		return "", errors.New("missing blockchain wallet address")
	}
	if rand.Float64() < 0.03 { // 模拟失败概率 3%
		return "", errors.New("blockchain payment failed")
	}
	transactionID := fmt.Sprintf("BC-%d", rand.Int63())
	return transactionID, nil
}

// Processing payment logic
func ProcessPayment(payment *models.Payment) (string, error) {
	var transactionID string
	var err error

	// Dynamically parse details based on payment method and call corresponding simulation service
	switch payment.Method {
	case "credit_card":
		var details CreditCardDetails
		if err := json.Unmarshal([]byte(payment.Details), &details); err != nil {
			return "", fmt.Errorf("failed to parse credit card details: %w", err)
		}
		transactionID, err = SimulateCreditCardPayment(details, payment.Amount)
	case "bank_transfer":
		var details BankTransferDetails
		if err := json.Unmarshal([]byte(payment.Details), &details); err != nil {
			return "", fmt.Errorf("failed to parse bank transfer details: %w", err)
		}
		transactionID, err = SimulateBankTransfer(details, payment.Amount)
	case "third_party":
		var details ThirdPartyDetails
		if err := json.Unmarshal([]byte(payment.Details), &details); err != nil {
			return "", fmt.Errorf("failed to parse third-party payment details: %w", err)
		}
		transactionID, err = SimulateThirdPartyPayment(details, payment.Amount)
	case "blockchain":
		var details BlockchainDetails
		if err := json.Unmarshal([]byte(payment.Details), &details); err != nil {
			return "", fmt.Errorf("failed to parse blockchain details: %w", err)
		}
		transactionID, err = SimulateBlockchainPayment(details, payment.Amount)
	default:
		err = errors.New("unsupported payment method")
	}

	return transactionID, err
}

// Create payment record
func CreatePayment(c *fiber.Ctx, db *gorm.DB) error {
	var payment models.Payment
	if err := c.BodyParser(&payment); err != nil {
		log.Printf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Process payments
	transactionID, err := ProcessPayment(&payment)
	if err != nil {
		log.Printf("Payment processing failed: %v", err)
		payment.Status = "failure"
		payment.TransactionID = ""
	} else {
		log.Printf("Payment processed successfully, transaction ID: %s", transactionID)
		payment.Status = "success"
		payment.TransactionID = transactionID
	}

	// save to database
	if err := db.Create(&payment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save payment"})
	}

	return c.Status(fiber.StatusCreated).JSON(payment)
}

// Get payment details
func GetPayment(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	var payment models.Payment
	if err := db.First(&payment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Payment not found"})
	}
	return c.JSON(payment)
}

// Register Payment Router
func RegisterPaymentRoutes(app fiber.Router, db *gorm.DB) {
	app.Post("payments/", func(c *fiber.Ctx) error {
		return CreatePayment(c, db)
	})
	app.Get("payments/:id", func(c *fiber.Ctx) error {
		return GetPayment(c, db)
	})
}
