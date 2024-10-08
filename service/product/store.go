package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/bagashyt/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetProductsByID implements types.ProductStore.
func (s *Store) GetProductsByID(productIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	// Convert ProductIDs to []interface{}
	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

// CreateProduct implements types.ProductStore.
func (s *Store) CreateProduct(product types.CreateProductPayload) error {
	_, err := s.db.Exec("INSERT INTO products (name, price, image, description, quantity) VALUE (?, ?, ?, ?, ?)", product.Name, product.Price, product.Image, product.Description, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

// GetProducts implements types.ProductStore.
func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	product := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		product = append(product, p)

	}

	return product, nil
}

// UpdateProduct implements types.ProductStore.
func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE INTO products (name, price, image, description, quantity) VALUES (?, ?, ?, ?, ?)", product.Name, product.Price, product.Image, product.Description, product.Quantity)
	if err != nil {
		return err
	}

	return nil
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
