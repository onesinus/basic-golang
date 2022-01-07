package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/onesinus/basic-golang/entities"
	"github.com/onesinus/basic-golang/user"
	"github.com/onesinus/basic-golang/auth"
)

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Simple Restful API")
	fmt.Println("Endpoint Hit: /")
}

func handleRequests() {
	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", welcomePage).Methods("GET")
	routes.HandleFunc("/users", user.Read).Methods("GET")
	routes.HandleFunc("/users/{id}", user.ReadOne).Methods("GET")
	routes.HandleFunc("/users", user.Create).Methods("POST")
	routes.HandleFunc("/users/{id}", user.Update).Methods("PUT")
	routes.HandleFunc("/users/{id}", user.Delete).Methods("DELETE")
	routes.HandleFunc("/auth/login", auth.Login).Methods("POST")

	fmt.Println("Server is running on port 7314")
	log.Fatal(http.ListenAndServe(":7314", routes))
}

func main() {
	entities.Users = []entities.User{
		entities.User{
			UserID: 1,
			Username: "onesinus",
			Email: "onesinus231@gmail.com",
			Address: "Jakarta",
			Password: "123",
		},
		entities.User{
			UserID: 2,
			Username: "jack",
			Email: "jack@gmail.com",
			Address: "Bandung",
			Password: "234",
		},
	}
	handleRequests()
}