package routes

import (
	"github.com/gorilla/mux"
	"github.com/snrweb/peer_to_peer_payment_app/config"
	"github.com/snrweb/peer_to_peer_payment_app/controllers"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc(config.API_VERSION_ONE+"user/add", controllers.CreateUser).Methods("POST", "OPTIONS")
	r.HandleFunc(config.API_VERSION_ONE+"user/all", controllers.GetUsers).Methods("GET", "OPTIONS")
	r.HandleFunc(config.API_VERSION_ONE+"user/{id}", controllers.GetUser).Methods("GET", "OPTIONS")
}
