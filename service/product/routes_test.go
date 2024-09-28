package product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bagashyt/ecom/types"
	"github.com/gorilla/mux"
)

func TestProductServiceHandlers(t *testing.T) {

	productStore := &mockProductStore{}
	userStore := &mockUserStore{}
	handler := NewHandler(productStore, userStore)

	t.Run("should handle get products", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/product", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleGetProducts).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}

type mockProductStore struct{}

// CreateProduct implements types.ProductStore.
func (m *mockProductStore) CreateProduct(types.CreateProductPayload) error {
	return nil
}

// GetProducts implements types.ProductStore.
func (m *mockProductStore) GetProducts() ([]*types.Product, error) {
	return []*types.Product{}, nil
}

// GetProductsByID implements types.ProductStore.
func (m *mockProductStore) GetProductsByID(ids []int) ([]types.Product, error) {
	return []types.Product{}, nil
}

// UpdateProduct implements types.ProductStore.
func (m *mockProductStore) UpdateProduct(types.Product) error {
	return nil
}

func (m *mockProductStore) GetProductByID(productID int) (*types.Product, error) {
	return &types.Product{}, nil
}

type mockUserStore struct{}

// CreateUser implements types.UserStore.
func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}

// GetUserByEmail implements types.UserStore.
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

// GetUserByID implements types.UserStore.
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}
