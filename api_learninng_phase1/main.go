package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

var users = make(map[int]User)
var nextID = 1

func createUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var user User

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    users[nextID] = user
    nextID++

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
func getUsers(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

 


func main() {
   
	 http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        createUser(w, r)
    } else if r.Method == http.MethodGet {
        getUsers(w, r)
    } else {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
})

 fmt.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}
