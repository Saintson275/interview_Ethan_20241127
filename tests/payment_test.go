package tests

import (
	"oxo/handlers"
	"oxo/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulateCreditCardPayment(t *testing.T) {
	// Simulate input data for credit card payment
	details := handlers.CreditCardDetails{
		CardNumber: "4111111111111111",
		ExpiryDate: "12/25",
		CVV:        "123",
	}
	amount := 100.0

	// Call the simulateCreditCardPayment function
	transactionID, err := handlers.SimulateCreditCardPayment(details, amount)

	// Check if success
	assert.Nil(t, err)
	assert.NotEmpty(t, transactionID)
}

func TestSimulateBankTransfer(t *testing.T) {
	// Simulate input data for bank transfer payment
	details := handlers.BankTransferDetails{
		AccountNumber: "1234567890",
		BankName:      "Barclay",
	}
	amount := 200.0

	// Call the simulateBankTransfer function
	transactionID, err := handlers.SimulateBankTransfer(details, amount)

	// Check if success
	assert.Nil(t, err)
	assert.NotEmpty(t, transactionID)
}

func TestSimulateThirdPartyPayment(t *testing.T) {
	// Simulate the input data of the third-party payment platform
	details := handlers.ThirdPartyDetails{
		Platform: "PayPal",
		Email:    "random@go.com",
	}
	amount := 50.0

	// Call simulateThirdPartyPayment
	transactionID, err := handlers.SimulateThirdPartyPayment(details, amount)

	// Check if success
	assert.Nil(t, err)
	assert.NotEmpty(t, transactionID)
}

func TestSimulateBlockchainPayment(t *testing.T) {
	// Input data for simulating blockchain payments
	details := handlers.BlockchainDetails{
		WalletAddress: "0x123abc456def789gh0",
		TransactionID: "T20741FGB612",
	}
	amount := 500.0

	// Call simulateBlockchainPayment
	transactionID, err := handlers.SimulateBlockchainPayment(details, amount)

	// Check if success
	assert.Nil(t, err)
	assert.NotEmpty(t, transactionID)
}

func TestProcessPaymentUnsupportedMethod(t *testing.T) {
	// Testing unsupported payment methods
	payment := &models.Payment{
		Method: "unsupported_method",
		Amount: 100.0,
	}

	transactionID, err := handlers.ProcessPayment(payment)

	// failed
	assert.NotNil(t, err)
	assert.Empty(t, transactionID)
}
