package services

import (
	"errors"
	"mooi/src/product/dto"
	"mooi/src/product/models"
	"mooi/src/product/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	AddProduct(request *dto.ProductRequest) error
	UpdateProduct(productID string, request *dto.ProductRequest) error
	DeleteProduct(productID string) error
}

type productService struct {
	ProductRepository repositories.ProductRepositoryInterface
}

func NewProductService(productRepository repositories.ProductRepositoryInterface) ProductService {
	return &productService{
		ProductRepository: productRepository,
	}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.ProductRepository.GetAllProducts()
}

func (s *productService) GetProductsByCategory(category string) ([]models.Product, error) {
	return s.ProductRepository.GetProductsByCategory(category)
}

func (s *productService) AddProduct(request *dto.ProductRequest) error {
	product := &models.Product{
		Name:     request.Name,
		Category: request.Category,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	return s.ProductRepository.AddProduct(product)
}

func (s *productService) UpdateProduct(productID string, request *dto.ProductRequest) error {
	product, err := s.ProductRepository.GetProductByID(productID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	product.Name = request.Name
	product.Category = request.Category
	product.Price = request.Price
	product.Quantity = request.Quantity

	return s.ProductRepository.UpdateProduct(product)
}

func (s *productService) DeleteProduct(productID string) error {
	product, err := s.ProductRepository.GetProductByID(productID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	return s.ProductRepository.DeleteProduct(productID)
}
