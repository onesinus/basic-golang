package auth

import (
	"fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/onesinus/basic-golang/entities"
)


func Login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: [POST] /auth/login")

    reqBody, _ := ioutil.ReadAll(r.Body)

    var auth entities.Auth
    var authResponse entities.AuthResponse

    json.Unmarshal(reqBody, &auth)

    isAuthorized := false
    for _, user := range entities.Users {
        if user.Password == auth.Password && user.Username == auth.Username {
			isAuthorized = true
        }
    }

    if isAuthorized == false {
    	w.WriteHeader(http.StatusUnauthorized)
    }

    authResponse.IsAuthorized = isAuthorized
    json.NewEncoder(w).Encode(authResponse)
}