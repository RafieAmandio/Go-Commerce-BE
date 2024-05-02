// handlers/cart_handler.go

package handlers

import (
	"encoding/json"
	"log"
	"mooi/library/response"
	"mooi/src/cart/dto"
	"mooi/src/cart/services"
	"net/http"

	"github.com/gorilla/mux"
)

type CartHandlers interface {
	GetCustomerCart(http.ResponseWriter, *http.Request)
	AddCartItem(http.ResponseWriter, *http.Request)
	RemoveCartItem(http.ResponseWriter, *http.Request)
}

type cartHandlers struct {
	CartService services.CartService
}

func NewCartHandler(cartService services.CartService) CartHandlers {
	return &cartHandlers{
		CartService: cartService,
	}
}

func (h *cartHandlers) GetCustomerCart(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	cart, err := h.CartService.GetCustomerCart(customerID)
	if err != nil {
		log.Println("Error getting customer cart:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to get customer cart", nil)
		return
	}

	response.JsonResponse(rw, http.StatusOK, "Customer cart retrieved successfully", cart)
}

func (h *cartHandlers) AddCartItem(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	productID := vars["product_id"]

	var request dto.AddCartItemRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error decoding request:", err)
		response.JsonResponse(rw, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	err = h.CartService.AddCartItem(customerID, productID, request.Quantity)
	if err != nil {
		log.Println("Error adding item to cart:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to add item to cart", nil)
		return
	}

	response.JsonResponse(rw, http.StatusCreated, "Item added to cart successfully", nil)
}

func (h *cartHandlers) RemoveCartItem(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// customerID := vars["customer_id"]
	cartItemID := vars["cart_item_id"]

	err := h.CartService.RemoveCartItem(cartItemID)
	if err != nil {
		log.Println("Error removing item from cart:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to remove item from cart", nil)
		return
	}

	response.JsonResponse(rw, http.StatusOK, "Item removed from cart successfully", nil)
}
