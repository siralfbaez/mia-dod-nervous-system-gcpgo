package validator

import (
	"errors"
	"fmt"
)

// OrderSignal matches the structure expected from the signal-gateway
type OrderSignal struct {
	OrderID  string  `json:"order_id"`
	Amount   float64 `json:"amount"`
	VendorID string  `json:"vendor_id"`
}

var (
	ErrMissingOrderID = errors.New("validation failed: missing order_id")
	ErrInvalidAmount  = errors.New("validation failed: amount must be greater than zero")
	ErrUnknownVendor  = errors.New("validation failed: unknown or unauthorized vendor")
)

// Validate checks the integrity of the incoming integration signal
func Validate(signal OrderSignal) error {
	if signal.OrderID == "" {
		return ErrMissingOrderID
	}
	if signal.Amount <= 0 {
		return ErrInvalidAmount
	}
	if signal.VendorID == "" {
		return ErrUnknownVendor
	}

	fmt.Printf("✅ Signal %s validated successfully for vendor %s\n", signal.OrderID, signal.VendorID)
	return nil
}