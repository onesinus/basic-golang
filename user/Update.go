package user

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "io/ioutil"
    "encoding/json"

    "github.com/onesinus/basic-golang/entities"
)


func Update(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: [PUT] /users")

    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        panic(err)
    }
    
    reqBody, _ := ioutil.ReadAll(r.Body)

    var updatedUser entities.User
    json.Unmarshal(reqBody, &updatedUser)

    for index, user := range entities.Users {
        if user.UserID == id {
            user.Username = updatedUser.Username
            user.Email = updatedUser.Email
            user.Address = updatedUser.Address
            user.Password = updatedUser.Password

            entities.Users[index] = user
            json.NewEncoder(w).Encode(user)
        }
    }
}