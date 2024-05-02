package dto

import (
	"mooi/src/order/models"
	"time"
)

// OrderRequest represents the request body for creating an order
type OrderRequest struct {
	CustomerID  string             `json:"customer_id"`
	TotalAmount float64            `json:"total_amount"`
	Items       []models.OrderItem `json:"items"`
}

// OrderResponse represents the response body for an order
type OrderResponse struct {
	OrderID     string
	ID          string             `json:"id"`
	CustomerID  string             `json:"customer_id"`
	TotalAmount float64            `json:"total_amount"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Items       []models.OrderItem `json:"items"`
}

type CreateOrderRequest struct {
	OrderID     string           `json:"order_id"`
	CustomerID  string           `json:"customer_id"`
	TotalAmount float64          `json:"total_amount"`
	Items       []OrderItemInput `json:"items"`
}

type OrderItemInput struct {
	OrderItemID string `json:"order_item_id"`
	ProductID   string `json:"product_id"`
	Quantity    int    `json:"quantity"`
}
