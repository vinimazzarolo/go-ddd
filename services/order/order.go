package order

import (
	"log"

	"github.com/google/uuid"
	"github.com/vinimazzarolo/go-ddd/domain/customer"
	custmem "github.com/vinimazzarolo/go-ddd/domain/customer/memory"
	"github.com/vinimazzarolo/go-ddd/domain/product"
	prodmem "github.com/vinimazzarolo/go-ddd/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := custmem.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}

		os.products = pr

		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIds []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []product.Product
	var total float64

	for _, id := range productIds {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return total, nil
}

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
