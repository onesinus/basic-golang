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
    w.Header().Set("Content-Type", "application/json")

    reqBody, _ := ioutil.ReadAll(r.Body)

    var user entities.User
    json.Unmarshal(reqBody, &user)

    user.UserID = len(entities.Users) + 1

    var ErrorMessage string

    if user.Username == "" {        
        ErrorMessage = "Username cannot be empty"
    }

    if user.Email == "" {
        ErrorMessage = "Email cannot be empty"
    }

    if user.Password == "" {
        ErrorMessage = "Password cannot be empty"
    }

    lengthPassword := len(user.Password)
    if lengthPassword < 6 || lengthPassword > 12 {
        ErrorMessage = "Password must between 6 - 12 characters"
    }

    if ErrorMessage != ""  {
        w.WriteHeader(http.StatusBadRequest)
        Response := entities.Response {
            Message: ErrorMessage,
        }

        json.NewEncoder(w).Encode(Response)
        return
    }
    
    entities.Users = append(entities.Users, user)
    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(user)
    
}