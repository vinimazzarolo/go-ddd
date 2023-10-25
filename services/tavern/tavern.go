package tavern

import (
	"log"

	"github.com/google/uuid"
	"github.com/vinimazzarolo/go-ddd/services/order"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	OrderService *order.OrderService

	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("\nBill the costumer: %0.0f\n", price)

	return nil
}

func WithBillingService(bs interface{}) TavernConfiguration {
	return func(t *Tavern) error {
		t.BillingService = bs
		return nil
	}
}
