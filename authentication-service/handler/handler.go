package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bernardn38/ebuy/authentication-service/models"
	"github.com/bernardn38/ebuy/authentication-service/service"
	"github.com/bernardn38/ebuy/authentication-service/token"
)

type Handler struct {
	authService  *service.AuthService
	tokenManager *token.Manager
}

func NewHandler(a *service.AuthService) *Handler {
	return &Handler{
		authService: a,
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	registerPayload := models.RegisterPayload{}
	err := json.NewDecoder(r.Body).Decode(&registerPayload)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to decode json body", http.StatusBadRequest)
		return
	}
	err = registerPayload.Validate()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
