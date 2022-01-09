# Run the application
1. install go on your operating system
2. execute **go run main.go** on your terminal

# API

| Endpoint    | Method | Example Request Body                              | Example Response                                    | Description    |
|-------------|--------|---------------------------------------------------|-----------------------------------------------------|----------------|
| /           | GET    |                                                   | Golang Simple Restful API                           | Root Endpoint  |
| /users      | GET    |                                                   | [{username: "bambang", email: "bambang@gmail.com"}] | Get users      |
| /users/{id} | GET    |                                                   | {username: "bambang", email: "bambang@gmail.com"}   | Get an user    |
| /users      | POST    | {username: "bambang", email: "bambang@gmail.com"} | {username: "bambang", email: "bambang@gmail.com"}   | Add an user |
| /users/{id} | PUT    | {username: "bambang", email: "bambang@gmail.com"} | {username: "bambang", email: "bambang@gmail.com"}   | Update an user |
| /users/{id} | DELETE |                                                   | {username: "bambang", email: "bambang@gmail.com"}   | Delete an user |
| /auth/login | POST   | {username: "onesinus", password: 123}             | {is_authorized: true}                                | Login          |
