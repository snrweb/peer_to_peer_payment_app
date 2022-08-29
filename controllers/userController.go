package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/snrweb/peer_to_peer_payment_app/cores"
	"github.com/snrweb/peer_to_peer_payment_app/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	err := cores.Validate(r.FormValue("first_name"), "First Name", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	err = cores.Validate(r.FormValue("last_name"), "Last Name", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	err = cores.Validate(r.FormValue("email"), "Email", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	u := &models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	user, err := u.CreateUser()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	userAccount := &models.Account{
		UserID: user.ID,
	}

	_, err = userAccount.Create()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{"status": true, "message": "User created successfully"}
	json.NewEncoder(w).Encode(data)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	u := &models.User{}

	user, err := u.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{"status": true, "data": user}
	json.NewEncoder(w).Encode(data)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	err := cores.Validate(vars["id"], "User ID", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	u := &models.User{
		ID: vars["id"],
	}

	user, err := u.GetUser()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{"status": true, "data": user}
	json.NewEncoder(w).Encode(data)
}

func updateUser() {
	// Update user
}

func deleteUser() {
	// delete user
}
