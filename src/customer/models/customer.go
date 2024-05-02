package models

type Customer struct {
    CustomerID   string `json:"customer_id"`
    Username     string `json:"username"`
    Email        string `json:"email"`
    Password     string `json:"password"`
    ProfilePhoto string `json:"profile_photo"`
    CreatedAt    string `json:"created_at"`
    UpdatedAt    string `json:"updated_at"`
}
