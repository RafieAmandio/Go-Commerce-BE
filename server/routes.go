package server

import (
	"mooi/library/middleware"
	customerHandlers "mooi/src/customer/handlers"
	customerRepositories "mooi/src/customer/repositories"
	customerServices "mooi/src/customer/services"
	productHandlers "mooi/src/product/handlers"
	productRepositories "mooi/src/product/repositories"
	productServices "mooi/src/product/services"
)

func (s *Server) routes() {

	// logging middleware which log total time taken to complete request
	s.Router.Use(middleware.LoggingMiddleware)

	// Initialize customer repository
	customerRepository := customerRepositories.NewCustomerRepository(s.Database)
	// Initialize customer service with the repository
	customerService := customerServices.NewCustomerService(customerRepository)
	// Initialize customer handlers with the service
	customerHandler := customerHandlers.NewCustomerHandler(customerService)

	// Define customer routes
	s.Router.HandleFunc("/register", customerHandler.RegisterCustomer).Methods("POST")
	s.Router.HandleFunc("/login", customerHandler.LoginCustomer).Methods("POST")

	// Initialize product repository
	productRepository := productRepositories.NewProductRepository(s.Database)
	// Initialize product service with the repository
	productService := productServices.NewProductService(productRepository)
	// Initialize product handlers with the service
	productHandler := productHandlers.NewProductHandler(productService)

	// Define product routes
	s.Router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	s.Router.HandleFunc("/products/{category}", productHandler.GetProductsByCategory).Methods("GET")
	s.Router.HandleFunc("/products", productHandler.AddProduct).Methods("POST")
	s.Router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	s.Router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	
}


