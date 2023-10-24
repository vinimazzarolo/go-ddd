package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/vinimazzarolo/go-ddd/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "IPA", 6.99)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := aggregate.NewProduct("Peenuts", "Snacks", 2.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "Red", 12.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{beer, peenuts, wine}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("John")
	if err != nil {
		t.Fatal(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}

}
