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

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}

	for _, user := range entities.Users {
		if user.UserID == id {
			json.NewEncoder(w).Encode(user)
		}
	}
}