package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bernardn38/ebuy/product-service/models"
	"github.com/bernardn38/ebuy/product-service/service"
	"github.com/bernardn38/ebuy/product-service/token"
)

type Handler struct {
	productService *service.ProductService
	tokenManager   *token.Manager
}

func NewHandler(p *service.ProductService, tm *token.Manager) *Handler {
	return &Handler{
		productService: p,
		tokenManager:   tm,
	}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductParams models.CreateProductParams
	json.NewDecoder(r.Body).Decode(&createProductParams)
	err := createProductParams.Validate()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	productId, err := h.productService.CreateProduct(r.Context(), createProductParams)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"productId": productId,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAllProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetRecentUploadedProducts(w http.ResponseWriter, r *http.Request) {
	recentProducts, err := h.productService.GetRecentUploadedProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNoContent)
	}
	json.NewEncoder(w).Encode(recentProducts)
	if err != nil {
		log.Println(err)
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}
