package dto

type AddCartItemRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CartItemResponse struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CartResponse struct {
	ID         string             `json:"id"`
	CustomerID string             `json:"customer_id"`
	CreatedAt  string             `json:"created_at"`
	UpdatedAt  string             `json:"updated_at"`
	Items      []CartItemResponse `json:"items"`
}
