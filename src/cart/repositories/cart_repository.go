// repositories/cart_repository.go

package repositories

import (
	"database/sql"
	"mooi/src/cart/models"
	"mooi/storage/db"
)

type CartRepositoryInterface interface {
	GetCartByCustomerID(customerID string) (*models.Cart, error)
	AddCartItem(cartID, productID string, quantity int) error
	RemoveCartItem(cartItemID string) error
}

type cartRepository struct {
	database db.PostgreSQLDB
}

func NewCartRepository(database db.PostgreSQLDB) CartRepositoryInterface {
	return &cartRepository{database: database}
}

func (r *cartRepository) GetCartByCustomerID(customerID string) (*models.Cart, error) {
	// Database connection
	db := r.database.ConnectDatabase()

	query := `
		SELECT cart_id, customer_id, created_at, updated_at
		FROM Cart
		WHERE customer_id = $1
	`
	row := db.QueryRow(query, customerID)

	var cart models.Cart
	err := row.Scan(&cart.ID, &cart.CustomerID, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Fetch cart items for the cart
	cart.Items, err = r.getCartItems(db, cart.ID)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *cartRepository) getCartItems(db *sql.DB, cartID string) ([]models.CartItem, error) {
	query := `
		SELECT cart_item_id, product_id, quantity, created_at, updated_at
		FROM CartItem
		WHERE cart_id = $1
	`
	rows, err := db.Query(query, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItem
	for rows.Next() {
		var item models.CartItem
		err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *cartRepository) AddCartItem(cartID, productID string, quantity int) error {
	// Database connection
	db := r.database.ConnectDatabase()

	query := `
		INSERT INTO CartItem (cart_item_id, cart_id, product_id, quantity)
		VALUES (gen_random_uuid(), $1, $2, $3)
	`
	_, err := db.Exec(query, cartID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (r *cartRepository) RemoveCartItem(cartItemID string) error {
	// Database connection
	db := r.database.ConnectDatabase()

	query := `
		DELETE FROM CartItem
		WHERE cart_item_id = $1
	`
	_, err := db.Exec(query, cartItemID)
	if err != nil {
		return err
	}
	return nil
}
