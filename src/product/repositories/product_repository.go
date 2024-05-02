// repositories/product_repository.go

package repositories

import (
	"database/sql"
	"errors"
	"mooi/src/product/models"
	"mooi/storage/db"
)

type ProductRepositoryInterface interface {
	GetAllProducts() ([]models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	GetProductByID(productID string) (*models.Product, error)
	AddProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID string) error
}

type productRepository struct {
	database db.PostgreSQLDB
}

func NewProductRepository(database db.PostgreSQLDB) ProductRepositoryInterface {
	return &productRepository{database: database}
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	query := `
        SELECT product_id, name, category, price, quantity
        FROM Product
    `

	rows, err := r.database.ConnectDatabase().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ProductID, &product.Name, &product.Category, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) GetProductsByCategory(category string) ([]models.Product, error) {
	query := `
        SELECT product_id, name, category, price, quantity
        FROM Product
        WHERE category = $1
    `

	rows, err := r.database.ConnectDatabase().Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ProductID, &product.Name, &product.Category, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) GetProductByID(productID string) (*models.Product, error) {
	query := `
        SELECT product_id, name, category, price, quantity
        FROM Product
        WHERE product_id = $1
    `

	row := r.database.ConnectDatabase().QueryRow(query, productID)

	var product models.Product
	err := row.Scan(&product.ProductID, &product.Name, &product.Category, &product.Price, &product.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no product found
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) AddProduct(product *models.Product) error {
	query := `
        INSERT INTO Product (product_id, name, category, price, quantity)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err := r.database.ConnectDatabase().Exec(query, product.ProductID, product.Name, product.Category, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	query := `
        UPDATE Product
        SET name = $2, category = $3, price = $4, quantity = $5
        WHERE product_id = $1
    `

	_, err := r.database.ConnectDatabase().Exec(query, product.ProductID, product.Name, product.Category, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProduct(productID string) error {
	query := `
        DELETE FROM Product
        WHERE product_id = $1
    `

	_, err := r.database.ConnectDatabase().Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}
