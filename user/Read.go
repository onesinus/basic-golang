package user

import (
	"fmt"
    "net/http"
    "encoding/json"

    "github.com/onesinus/basic-golang/entities"
)

func Read(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: [GET] /users")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(entities.Users)
}