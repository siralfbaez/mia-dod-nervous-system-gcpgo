package validator

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		signal  OrderSignal
		wantErr error
	}{
		{"Valid Signal", OrderSignal{OrderID: "123", Amount: 10.50, VendorID: "PH-01"}, nil},
		{"Missing Order ID", OrderSignal{Amount: 10.50, VendorID: "PH-01"}, ErrMissingOrderID},
		{"Negative Amount", OrderSignal{OrderID: "123", Amount: -1.0, VendorID: "PH-01"}, ErrInvalidAmount},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.signal); err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}