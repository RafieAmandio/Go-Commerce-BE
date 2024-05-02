package server

import (
	"mooi/library/middleware"
	"mooi/src/customer/handlers"
	"mooi/src/customer/repositories"
	"mooi/src/customer/services"
)

func (s *Server) routes() {

	// logging middleware which log total time taken to complete request
	s.Router.Use(middleware.LoggingMiddleware)

	// Initialize customer repository
	customerRepository := repositories.NewCustomerRepository(s.Database)
	// Initialize customer service with the repository
	customerService := services.NewCustomerService(customerRepository)
	// Initialize customer handlers with the service
	customerHandlers := handlers.NewCustomerHandler(customerService)

	// Define customer routes
	s.Router.HandleFunc("/register", customerHandlers.RegisterCustomer).Methods("POST")
	s.Router.HandleFunc("/login", customerHandlers.LoginCustomer).Methods("POST")

	// // define extra routes
	// exampleRepository := repositories.NewExampleRepostiory(s.Database)
	// exampleService := services.NewExampleService(exampleRepository)
	// exampleHandlers := handlers.NewHttpHandler(exampleService)
	// // setup example routes
	// s.Router.HandleFunc("/", exampleHandlers.GetExample)
	// s.Router.HandleFunc("/test", exampleHandlers.GetExample)
}
