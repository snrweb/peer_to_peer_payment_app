package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/snrweb/peer_to_peer_payment_app/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	r := mux.NewRouter()

	//----------
	// Routes
	//----------
	routes.UserRoutes(r)
	routes.AccountRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
