package register

import (
	"encoding/json"
	"log"
	"net/http"
	"parking/internal/register"

	"github.com/gorilla/mux"
)

type Handler struct {
	UseCase register.UseCase
}

func NewHandler(useCase register.UseCase) Handler {
	return Handler{useCase}
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq register.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print("Create User Request: ", userReq)
	err = h.UseCase.CreateUser(r.Context(), userReq)
	if err != nil {
		errMessage := errorMessage{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(errMessage)
		if err != nil {
			log.Printf("Error encoding error message: %v", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	user, err := h.UseCase.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userReq register.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print("Update User Request: ", userReq)
	err = h.UseCase.UpdateUser(r.Context(), userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UseCase.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

type errorMessage struct {
	Message string `json:"message"`
}
