package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json: "Age"`
}

func profileHandler(w http.ResponseWriter,r * http.Request){
	w.Header().Set("Content-Type","application/json")
	user:=User{
		Name:"Ganesh",
		Email:"ganeshh9876@gmail.com",
		Age:22,
		
	}
	json.NewEncoder(w).Encode(user)

	

}

func createprofilehandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	if r.Method!=http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error":"only post allowed"}`) )
		return 
	}
  var u User
  err:=json.NewDecoder(r.Body).Decode(&u)
  if err!=nil{
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error":"Invalid Json"}`))
	return
  }
  json.NewEncoder(w).Encode(u)
}

func main() {
   http.HandleFunc("/profile",profileHandler)
  http.HandleFunc("/create",createprofilehandler)
	http.ListenAndServe(":8080",nil)

}
