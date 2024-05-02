package services

import (
	"mooi/src/cart/dto"
	"mooi/src/cart/repositories"
	"time"
)

type CartService interface {
	GetCustomerCart(customerID string) (*dto.CartResponse, error)
	AddCartItem(customerID, productID string, quantity int) error
	RemoveCartItem(cartItemID string) error
}

type cartService struct {
	CartRepository repositories.CartRepositoryInterface
}

func NewCartService(cartRepository repositories.CartRepositoryInterface) CartService {
	return &cartService{
		CartRepository: cartRepository,
	}
}

func (s *cartService) GetCustomerCart(customerID string) (*dto.CartResponse, error) {
	// Retrieve cart from repository
	cart, err := s.CartRepository.GetCartByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	var cartItemsResponse []dto.CartItemResponse
	for _, item := range cart.Items {
		cartItemResponse := dto.CartItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		cartItemsResponse = append(cartItemsResponse, cartItemResponse)
	}

	// Convert cart to DTO
	response := dto.CartResponse{
		ID:         cart.ID,
		CustomerID: cart.CustomerID,
		CreatedAt:  cart.CreatedAt.Format(time.RFC3339), // Typecast to string
		UpdatedAt:  cart.UpdatedAt.Format(time.RFC3339), // Typecast to string
		Items:      cartItemsResponse,
	}

	return &response, nil
}

func (s *cartService) AddCartItem(customerID, productID string, quantity int) error {
	// Get cart for customer
	cart, err := s.CartRepository.GetCartByCustomerID(customerID)
	if err != nil {
		return err
	}

	// Add item to cart
	err = s.CartRepository.AddCartItem(cart.ID, productID, quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *cartService) RemoveCartItem(cartItemID string) error {
	// Remove item from cart
	err := s.CartRepository.RemoveCartItem(cartItemID)
	if err != nil {
		return err
	}

	return nil
}
