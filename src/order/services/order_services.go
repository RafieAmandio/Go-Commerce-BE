package services

import (
	"mooi/src/order/dto"
	"mooi/src/order/models"
	"mooi/src/order/repositories"
)

type OrderService interface {
	CreateOrder(request *dto.CreateOrderRequest) error
	GetCustomerOrders(customerID string) ([]*dto.OrderResponse, error)
}

type orderService struct {
	orderRepository repositories.OrderRepositoryInterface
}

func NewOrderService(orderRepository repositories.OrderRepositoryInterface) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (s *orderService) CreateOrder(request *dto.CreateOrderRequest) error {
	// Create the order
	order := &models.Order{
		OrderID:     request.OrderID,
		CustomerID:  request.CustomerID,
		TotalAmount: request.TotalAmount,
	}
	err := s.orderRepository.CreateOrder(order)
	if err != nil {
		return err
	}

	// Create order items
	for _, item := range request.Items {
		orderItem := &models.OrderItem{
			OrderItemID: item.OrderItemID,
			OrderID:     request.OrderID,
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
		}
		err := s.orderRepository.CreateOrderItem(orderItem)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *orderService) GetCustomerOrders(customerID string) ([]*dto.OrderResponse, error) {
	orders, err := s.orderRepository.GetOrdersByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	orderResponses := make([]*dto.OrderResponse, 0, len(orders))
	for _, order := range orders {
		orderResponse := &dto.OrderResponse{
			OrderID:     order.OrderID,
			CustomerID:  order.CustomerID,
			TotalAmount: order.TotalAmount,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		}
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}
