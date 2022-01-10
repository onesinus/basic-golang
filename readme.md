# Deployed Application (for testing purpose)
http://ec2-13-214-161-205.ap-southeast-1.compute.amazonaws.com:7314/

# Run the application
1. install **go** on your operating system
2. Fork this repository, then **clone the forked repository to your computer**
3. **Open terminal and change directory location** to the project
4. execute **go run main.go** on your terminal

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
