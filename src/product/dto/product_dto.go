
package dto

type ProductRequest struct {
    Name     string  `json:"name"`
    Category string  `json:"category"`
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
}

type ProductResponse struct {
    ProductID string  `json:"product_id"`
    Name      string  `json:"name"`
    Category  string  `json:"category"`
    Price     float64 `json:"price"`
    Quantity  int     `json:"quantity"`
}
