package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/snrweb/peer_to_peer_payment_app/cores"
	"github.com/snrweb/peer_to_peer_payment_app/models"
)

type myAccount models.Account

func (account myAccount) Create() (myAccount, error) {
	err := cores.Validate(account.UserID, "User ID", []string{"empty"})
	if err != nil {
		return account, err
	}

	account.Balance = 0.0

	account, err = account.Create()
	if err != nil {
		return account, err
	}

	return account, nil
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	err := cores.Validate(vars["user_id"], "User ID", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	account := &models.Account{
		UserID: vars["user_id"],
	}

	result, err := account.GetBalance()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{"status": true, "data": map[string]float64{"balance": result}}
	json.NewEncoder(w).Encode(data)
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	err := cores.Validate(r.FormValue("user_id"), "User ID", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)

	err = cores.Validate(amount, "Amount", []string{"greaterThanZero"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	account := &models.Account{
		UserID:  r.FormValue("user_id"),
		Balance: amount,
	}

	result, err := account.Deposit()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{"status": true, "data": result}
	json.NewEncoder(w).Encode(data)
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	err := cores.Validate(r.FormValue("user_id"), "User ID", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
	}

	err = cores.Validate(r.FormValue("recipient"), "Recipient ID", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)

	err = cores.Validate(amount, "Amount", []string{"greaterThanZero"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	u := &models.User{
		ID: r.FormValue("recipient"),
	}

	account := &models.Account{
		UserID:  r.FormValue("user_id"),
		Balance: amount,
	}

	result, err := account.Transfer(*u)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{"status": true, "data": result}
	json.NewEncoder(w).Encode(data)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	err := cores.Validate(r.FormValue("user_id"), "User ID", []string{"empty"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)

	err = cores.Validate(amount, "Amount", []string{"greaterThanZero"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	account := &models.Account{
		UserID:  r.FormValue("user_id"),
		Balance: amount,
	}

	result, err := account.Withdraw()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		data := map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{"status": true, "data": result}
	json.NewEncoder(w).Encode(data)
}
