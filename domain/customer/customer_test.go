package customer_test

import (
	"errors"
	"testing"

	"github.com/vinimazzarolo/go-ddd/domain/customer"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: customer.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "John Doe",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := customer.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error to be %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
