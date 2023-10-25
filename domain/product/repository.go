package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductAlreadyExists = errors.New("the product already exists")
	ErrProductNotFound      = errors.New("the product was not found")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(p Product) error
	Update(p Product) error
	Delete(id uuid.UUID) error
}
