package user

import (
	"fmt"
    "net/http"
    "encoding/json"
	"github.com/gorilla/mux"
	"strconv"

    "github.com/onesinus/basic-golang/entities"
)

func ReadOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: [GET] /users/{UserId}")

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}

	isFound := false

	for _, user := range entities.Users {
		if user.UserID == id {
			isFound = true
			json.NewEncoder(w).Encode(user)
		}
	}

	if isFound == false {
		w.WriteHeader(http.StatusNotFound)
		Response := entities.Response {
			Message: "User tidak ditemukan",
		}

		json.NewEncoder(w).Encode(Response)	
	}
}