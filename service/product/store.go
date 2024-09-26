package product

import (
	"database/sql"
	"fmt"

	"github.com/bagashyt/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetProductsByID implements types.ProductStore.
func (s *Store) GetProductsByID(ids []int) ([]types.Product, error) {
	panic("unimplemented")
}

// CreateProduct implements types.ProductStore.
func (s *Store) CreateProduct(types.CreateProductPayload) error {
	panic("unimplemented")
}

// GetProducts implements types.ProductStore.
func (s *Store) GetProducts() ([]*types.Product, error) {
	panic("unimplemented")
}

// UpdateProduct implements types.ProductStore.
func (s *Store) UpdateProduct(types.Product) error {
	panic("unimplemented")
}

func (s *Store) GetProductByID(productID int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", productID)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, fmt.Errorf("Product not found")
	}
	return p, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
