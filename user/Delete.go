package user

import (
	"fmt"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "encoding/json"

    "github.com/onesinus/basic-golang/entities"
)


func Delete(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: [DELETE] /users")

    w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        panic(err)
    }
    
    isFound := false

    for index, user := range entities.Users {
        if user.UserID == id {
            isFound = true

            entities.Users = append(entities.Users[:index], entities.Users[index+1:]...)
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