package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bernardn38/ebuy/authentication-service/models"
	"github.com/bernardn38/ebuy/authentication-service/service"
	"github.com/bernardn38/ebuy/authentication-service/token"
	"github.com/lib/pq"
)

type Handler struct {
	authService  *service.AuthService
	tokenManager *token.Manager
}

func NewHandler(a *service.AuthService, tm token.Manager) *Handler {
	return &Handler{
		authService:  a,
		tokenManager: &tm,
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
	err = h.authService.RegisterUser(r.Context(), registerPayload)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			// Handle Postgres-specific error
			switch pgErr.Code.Name() {
			case "unique_violation":
				// Handle unique constraint violation error
				http.Error(w, "username or email already taken", http.StatusBadRequest)
				return
			}
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	loginPayload := models.LoginPayload{}
	err := json.NewDecoder(r.Body).Decode(&loginPayload)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to decode json body", http.StatusBadRequest)
		return
	}
	err = loginPayload.Validate()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.authService.LoginUser(r.Context(), loginPayload)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			// Handle Postgres-specific error
			switch pgErr.Code.Name() {
			case "unique_violation":
				// Handle unique constraint violation error
				http.Error(w, "username or email already taken", http.StatusBadRequest)
				return
			}
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.tokenManager.GenerateToken(string(user.ID), user.Username, time.Hour)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "jwtToken",
		Domain:   "localhost",
		Path:     "/",
		Value:    token.String(),
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(map[string]int32{
		"userId": user.ID,
	})
	w.WriteHeader(http.StatusOK)
}
