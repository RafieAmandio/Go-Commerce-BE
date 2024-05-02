package repositories

import (
	"mooi/src/customer/models"
	"mooi/storage/db"
)

type CustomerRepositoryInterface interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomerByUsername(username string) (*models.Customer, error)
}

type customerRepository struct {
	database db.PostgreSQLDB
}

func NewCustomerRepository(database db.PostgreSQLDB) CustomerRepositoryInterface {
	return &customerRepository{database: database}
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) error {
	// Prepare SQL statement
	query := `
        INSERT INTO Customer (customer_id, username, email, password, profile_photo, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	// Execute SQL statement
	_, err := r.database.ConnectDatabase().Exec(query, customer.CustomerID, customer.Username, customer.Email, customer.Password, customer.ProfilePhoto, customer.CreatedAt, customer.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) GetCustomerByUsername(username string) (*models.Customer, error) {
	// Prepare SQL statement
	query := `
        SELECT customer_id, username, email, password, profile_photo, created_at, updated_at
        FROM Customer
        WHERE username = $1
    `
	// Execute SQL query
	row := r.database.ConnectDatabase().QueryRow(query, username)

	// Scan the result into a Customer struct
	var customer models.Customer
	err := row.Scan(&customer.CustomerID, &customer.Username, &customer.Email, &customer.Password, &customer.ProfilePhoto, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
