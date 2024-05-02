package services

import (
	"errors"
	"mooi/library/hash"
	"mooi/library/jwt"
	"mooi/src/customer/dto"
	"mooi/src/customer/models"
	"mooi/src/customer/repositories"
)

type CustomerService interface {
	RegisterCustomer(request *dto.CustomerRegisterRequest) error
	LoginCustomer(request *dto.CustomerLoginRequest) (*dto.CustomerLoginResponse, error)
}

type customerService struct {
	CustomerRepository repositories.CustomerRepositoryInterface
}

func NewCustomerService(customerRepository repositories.CustomerRepositoryInterface) CustomerService {
	return &customerService{
		CustomerRepository: customerRepository,
	}
}

func (s *customerService) RegisterCustomer(request *dto.CustomerRegisterRequest) error {
	// Hash the password

	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		return err
	}

	// Pass hashed password to repository for registration
	customer := &models.Customer{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
		// other fields can be set if needed
	}

	err = s.CustomerRepository.CreateCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}

func (s *customerService) LoginCustomer(request *dto.CustomerLoginRequest) (*dto.CustomerLoginResponse, error) {
	// Retrieve customer by username from repository
	customer, err := s.CustomerRepository.GetCustomerByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	// Compare hashed password
	err = hash.CompareHashAndPassword(customer.Password, request.Password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := jwt.GenerateJWTToken(customer.CustomerID)
	if err != nil {
		return nil, err
	}

	response := &dto.CustomerLoginResponse{
		Token: token,
	}

	return response, nil
}

// Example error definition for invalid credentials
var ErrInvalidCredentials = errors.New("invalid username or password")
