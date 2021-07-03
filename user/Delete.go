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

	vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        panic(err)
    }
    

    for index, user := range entities.Users {
        if user.UserID == id {
            entities.Users = append(entities.Users[:index], entities.Users[index+1:]...)
            json.NewEncoder(w).Encode(user)
        }
    }
}