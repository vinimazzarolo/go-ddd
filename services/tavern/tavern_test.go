package tavern

import (
	"testing"

	"github.com/google/uuid"
	"github.com/vinimazzarolo/go-ddd/domain/product"
	"github.com/vinimazzarolo/go-ddd/services/order"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "IPA", 6.99)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := product.NewProduct("Peenuts", "Snacks", 2.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "Red", 12.99)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{beer, peenuts, wine}
}

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("John")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}
}
