package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/vinimazzarolo/go-ddd/aggregate"
)

var (
	ErrProductAlreadyExists = errors.New("the product already exists")
	ErrProductNotFound      = errors.New("the product was not found")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(p aggregate.Product) error
	Update(p aggregate.Product) error
	Delete(id uuid.UUID) error
}
