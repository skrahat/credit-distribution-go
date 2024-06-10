package handlers

import (
	"encoding/json"
	"net/http"

	"credit-distribution-go/internal/services"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
    r.HandleFunc("/users", getAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", getUser).Methods("GET")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := services.GetAllUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    user, err := services.GetUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}
