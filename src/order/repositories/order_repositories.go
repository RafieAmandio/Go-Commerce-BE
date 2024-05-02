package repositories

import (
	"database/sql"
	"mooi/src/order/models"
)

type OrderRepositoryInterface interface {
	CreateOrder(order *models.Order) error
	GetOrdersByCustomerID(customerID string) ([]*models.Order, error)
	CreateOrderItem(orderItem *models.OrderItem) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepositoryInterface {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	query := `
		INSERT INTO "Order" (order_id, customer_id, total_amount)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, order.OrderID, order.CustomerID, order.TotalAmount)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) GetOrdersByCustomerID(customerID string) ([]*models.Order, error) {
	query := `
		SELECT order_id, customer_id, total_amount, created_at, updated_at
		FROM "Order"
		WHERE customer_id = $1
		
	`

	rows, err := r.db.Query(query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []*models.Order{}
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.OrderID, &order.CustomerID, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) CreateOrderItem(orderItem *models.OrderItem) error {
	query := `
		INSERT INTO OrderItem (order_item_id, order_id, product_id, quantity)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(query, orderItem.OrderItemID, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity)
	if err != nil {
		return err
	}

	return nil
}
