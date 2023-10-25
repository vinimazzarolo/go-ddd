package main

import (
	"github.com/google/uuid"
	"github.com/vinimazzarolo/go-ddd/domain/product"
	"github.com/vinimazzarolo/go-ddd/services/order"
	"github.com/vinimazzarolo/go-ddd/services/tavern"
)

func main() {
	products := productInventory()
	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("John")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "IPA", 6.99)
	if err != nil {
		panic(err)
	}

	peenuts, err := product.NewProduct("Peenuts", "Snacks", 2.99)
	if err != nil {
		panic(err)
	}

	wine, err := product.NewProduct("Wine", "Red", 12.99)
	if err != nil {
		panic(err)
	}

	return []product.Product{beer, peenuts, wine}
}
