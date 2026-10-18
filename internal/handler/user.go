package handler

import (
	"encoding/json"
	"net/http"

	"github.com/graimond/cazino/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

type RegisterRequest struct {
	Name string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Name string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler)Register(w http.ResponseWriter, r *http.Request){
	var req RegisterRequest

	err:=json.NewDecoder(r.Body).Decode(&req)

	if err!=nil{
		http.Error(w,"ivalid request body",http.StatusBadRequest)
		return
	}

	user,err:=h.service.Register(
		r.Context(),
		req.Name,
		req.Email,
		req.Password,
	)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Login (w http.ResponseWriter, r *http.Request){
	var req LoginRequest

	err:=json.NewDecoder(r.Body).Decode(&req)

	token,err:=h.service.Login(
		r.Context(),
		req.Name,
		req.Email,
		req.Password,
	)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(
		map[string]string{
			"token":token,
		})
}