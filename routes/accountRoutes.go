package routes

import (
	"github.com/gorilla/mux"
	"github.com/snrweb/peer_to_peer_payment_app/config"
	"github.com/snrweb/peer_to_peer_payment_app/controllers"
)

func AccountRoutes(r *mux.Router) {
	r.HandleFunc(config.API_VERSION_ONE+"account/deposit", controllers.Deposit).Methods("POST", "OPTIONS")
	r.HandleFunc(config.API_VERSION_ONE+"account/transfer", controllers.Transfer).Methods("POST", "OPTIONS")
	r.HandleFunc(config.API_VERSION_ONE+"account/withdraw", controllers.Withdraw).Methods("POST", "OPTIONS")
	r.HandleFunc(config.API_VERSION_ONE+"account/balance/{user_id}", controllers.GetBalance).Methods("GET", "OPTIONS")
}
