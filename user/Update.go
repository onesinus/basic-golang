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
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        panic(err)
    }
    
    reqBody, _ := ioutil.ReadAll(r.Body)

    var updatedUser entities.User
    json.Unmarshal(reqBody, &updatedUser)


    var ErrorMessage string

    if updatedUser.Username == "" {        
        ErrorMessage = "Username cannot be empty"
    }

    if updatedUser.Email == "" {
        ErrorMessage = "Email cannot be empty"
    }

    if updatedUser.Password == "" {
        ErrorMessage = "Password cannot be empty"
    }

    lengthPassword := len(updatedUser.Password)
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

    isFound := false

    for index, user := range entities.Users {
        if user.UserID == id {
            isFound = true

            user.Username = updatedUser.Username
            user.Email = updatedUser.Email
            user.Address = updatedUser.Address
            user.Password = updatedUser.Password

            entities.Users[index] = user
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