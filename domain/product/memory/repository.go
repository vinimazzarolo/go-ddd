package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/vinimazzarolo/go-ddd/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product

	for _, p := range mpr.products {
		products = append(products, p)
	}

	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if p, ok := mpr.products[id]; ok {
		return p, nil
	}

	return product.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(p product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	mpr.products[p.GetID()] = p

	return nil
}

func (mpr *MemoryProductRepository) Update(p product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[p.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[p.GetID()] = p

	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mpr.products, id)

	return nil
}
