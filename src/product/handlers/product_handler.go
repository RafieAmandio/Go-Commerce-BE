package handlers

import (
	"log"
	"mooi/library/response"
	"mooi/src/product/dto"
	"mooi/src/product/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandlers interface {
	GetAllProducts(http.ResponseWriter, *http.Request)
	GetProductsByCategory(http.ResponseWriter, *http.Request)
	AddProduct(http.ResponseWriter, *http.Request)
	UpdateProduct(http.ResponseWriter, *http.Request)
	DeleteProduct(http.ResponseWriter, *http.Request)
}

type productHandlers struct {
	ProductService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandlers {
	return &productHandlers{
		ProductService: productService,
	}
}

func (h *productHandlers) GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		log.Println("Error fetching all products:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to fetch products", nil)
		return
	}

	response.JsonResponse(rw, http.StatusOK, "Success", products)
}

func (h *productHandlers) GetProductsByCategory(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := params["category"]

	products, err := h.ProductService.GetProductsByCategory(category)
	if err != nil {
		log.Println("Error fetching products by category:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to fetch products by category", nil)
		return
	}

	response.JsonResponse(rw, http.StatusOK, "Success", products)
}

func (h *productHandlers) AddProduct(rw http.ResponseWriter, r *http.Request) {
	var request dto.ProductRequest

	err := r.ParseMultipartForm(32 * 1024 * 1024)
	if err != nil {
		log.Println("Error parsing multipart form:", err)
		response.JsonResponse(rw, http.StatusBadRequest, "Error parsing multipart form", nil)
		return
	}

	request.Name = r.FormValue("name")
	request.Category = r.FormValue("category")
	request.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
	request.Quantity, err = strconv.Atoi(r.FormValue("quantity"))

	err = h.ProductService.AddProduct(&request)
	if err != nil {
		log.Println("Error adding product:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to add product", nil)
		return
	}

	response.JsonResponse(rw, http.StatusCreated, "Product added successfully", nil)
}

func (h *productHandlers) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	var request dto.ProductRequest

	err := r.ParseMultipartForm(32 * 1024 * 1024)
	if err != nil {
		log.Println("Error parsing multipart form:", err)
		response.JsonResponse(rw, http.StatusBadRequest, "Error parsing multipart form", nil)
		return
	}

	request.Name = r.FormValue("name")
	request.Category = r.FormValue("category")
	request.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
	request.Quantity, err = strconv.Atoi(r.FormValue("quantity"))

	err = h.ProductService.UpdateProduct(productID, &request)
	if err != nil {
		log.Println("Error updating product:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to update product", nil)
		return
	}

	response.JsonResponse(rw, http.StatusOK, "Product updated successfully", nil)
}

func (h *productHandlers) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	err := h.ProductService.DeleteProduct(productID)
	if err != nil {
		log.Println("Error deleting product:", err)
		response.JsonResponse(rw, http.StatusInternalServerError, "Failed to delete product", nil)
		return
	}

	response.JsonResponse(rw, http.StatusOK, "Product deleted successfully", nil)
}
