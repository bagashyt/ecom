package cart

import (
	"fmt"

	"github.com/bagashyt/ecom/types"
)

func getCartItemsIDs(items []types.CartCheckoutItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("ivalid quantity for product %d", item.ProductID)
		}
		productIds[i] = item.ProductID
	}

	return productIds, nil
}

func checkIfCartIsInStock(cartItems []types.CartCheckoutItem, product map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := product[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your card", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartITems []types.CartCheckoutItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartITems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
