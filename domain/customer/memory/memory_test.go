package memory

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/vinimazzarolo/go-ddd/aggregate"
	"github.com/vinimazzarolo/go-ddd/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("John Doe")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			test:        "Customer not found",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			test:        "Customer found",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error to be %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
