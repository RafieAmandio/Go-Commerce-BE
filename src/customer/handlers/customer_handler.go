package handlers

import (
	"encoding/json"
	"log"
	"mooi/library/response"
	"mooi/src/customer/dto"
	"mooi/src/customer/services"
	"net/http"
)

type CustomerHandlers interface {
	RegisterCustomer(http.ResponseWriter, *http.Request)
	LoginCustomer(http.ResponseWriter, *http.Request)
}

type customerHandlers struct {
	CustomerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) CustomerHandlers {
	return &customerHandlers{
		CustomerService: customerService,
	}
}

func (h *customerHandlers) RegisterCustomer(rw http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(32 * 1024 * 1024)
	if err != nil {
		log.Println("Error parsing multipart form:", err)
		response.JsonResponse(rw, http.StatusBadRequest, "Failed to parse multipart form", nil)
		return
	}

	// Access the form values
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validate form values (e.g., check for required fields)

	// Create CustomerRegisterRequest from form values
	request := dto.CustomerRegisterRequest{
		Username: username,
		Email:    email,
		Password: password,
	}
	// Call the CustomerService to register the customer
	err = h.CustomerService.RegisterCustomer(&request)
	if err != nil {
		log.Println("Error registering customer:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to register customer", nil)
		return
	}

	// Respond with success message
	response.JsonResponse(rw, http.StatusCreated, "Customer registered successfully", nil)
}

func (h *customerHandlers) LoginCustomer(rw http.ResponseWriter, r *http.Request) {
	var request dto.CustomerLoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error decoding request:", err)
		response.JsonResponse(rw, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	result, err := h.CustomerService.LoginCustomer(&request)
	if err != nil {
		log.Println("Error logging in customer:", err)
		response.JsonResponse(rw, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}

	response.JsonResponse(rw, http.StatusOK, "Login successful", result)
}
