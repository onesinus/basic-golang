package user

import (
	"fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/onesinus/basic-golang/entities"
)


func Create(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: [POST] /users")

    reqBody, _ := ioutil.ReadAll(r.Body)

    var user entities.User
    json.Unmarshal(reqBody, &user)

    entities.Users = append(entities.Users, user)
    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(user)
}