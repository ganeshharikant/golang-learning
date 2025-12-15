// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// // func hellohandler(w http.ResponseWriter,r* http.Request){
// // 	w.Header().Set("Content-Type", "application/json")

// // 	 w.Write([]byte(`{"name":"ganesh"}`))
// // }

// // type User struct {
// // 	Name  string `json:"name"`
// // 	Email string `json:"email"`
// // }
// // func createUserHandler(w http.ResponseWriter, r *http.Request) {
// // 	var user User

// // 	json.NewDecoder(r.Body).Decode(&user)

// // 	fmt.Println("Received user:", user.Name)

// // 	w.Write([]byte("User received: " + user.Name))
// // }

// type Profile struct {
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	Language string `json:"language"`
// }

// func profileHandler(w http.ResponseWriter, r* http.Request ){
//   w.Header().Set("Content-Type","application/json")
//   profile:=Profile{
// 	       Name:"Ganesh",
// 		   Age:10,
// 		   Language:"hindi",

//   }
// json.NewEncoder(w).Encode(profile)

// }

// func main() {
// 	// fmt.Println("hello server has been started on 8080")
// 	// http.ListenAndServe(":8080" ,nil)

// 	//  http.HandleFunc("/hello",hellohandler)

// 	 fmt.Println("server tsrated")
// 	  http.HandleFunc("/profile",profileHandler)

// 	  http.ListenAndServe(":8080", nil)

// 	//  http.HandleFunc("/user", createUserHandler)

// }

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SignupRequest struct{
	Name string   `json:"name"`
    Email string  `json:"email"` 	 
	
}
type SignupResponse struct{
	Message string   `json:"message"`
    User    SignupRequest `json:"user"` 	 
	
}

var users=make(map[string]SignupRequest)


func senddetails(w http.ResponseWriter ,r *http.Request){
    w.Header().Set("content-type","application/json")
if(r.Method!=http.MethodPost){
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
	return 
}

	var user SignupRequest
	
err:=json.NewDecoder(r.Body).Decode(&user)
    if err!=nil{
		w.Write([]byte("error in decoding the method"))
		return 
	}
	
	users[user.Email]=user
resp:=SignupResponse{
	Message: "thanks for registering",
	User: SignupRequest{
		   Name:user.Name,
		   Email:user.Email,   
	} ,
}


     
   




	json.NewEncoder(w).Encode(resp)


	


  

}

func getdetails(w http.ResponseWriter,r *http.Request){
w.Header().Set("Content-Type", "application/json")
	List:=[]SignupRequest{}

    for _,user:=range users{
		List = append(List, user)
	}
json.NewEncoder(w).Encode(List)





}


func updatedetails(w http.ResponseWriter,r *http.Request){
   
if r.Method!=http.MethodPut{
    w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("hey this method is not allowed"))
	return 
}

email:=r.URL.Path[len("/users/"):]

  fmt.Println("Requested email:", email)

exuser,ok:=users[email]

if !ok{
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("hey there is no user with these names"))
	return 

}

var udate SignupRequest
json.NewDecoder(r.Body).Decode(&udate)

if udate.Name!=exuser.Name{
	exuser.Name=udate.Name

}
users[email]=exuser

w.Header().Set("Content-type","application/json")

json.NewEncoder(w).Encode(exuser)











}

func deletedetails(w http.ResponseWriter,r *http.Request){
     
  if r.Method!=http.MethodDelete{
	 w.WriteHeader(http.StatusNotFound)
	 w.Write([]byte("invalid method"))
	 return
  }

  email:=r.URL.Path[len("/rm/"):]

  _,ok:=users[email]
  if !ok{
	fmt.Println("user not found")
	w.Write([]byte("user not found"))
	return 
  }
      

  delete(users,email)
 

w.Header().Set("Content-Type","application/json")

json.NewEncoder(w).Encode(map[string]string{
   "message":"User deleted Suceesfully",
    "email" : email,


})











}

func getsingle(w http.ResponseWriter,r * http.Request){
       if r.Method!=http.MethodGet{
		  w.WriteHeader(http.StatusMethodNotAllowed)
		  w.Write([]byte("method name is diffrent"))
           return 
	   }


	   email:=r.URL.Path[len("/user/"):]
         
	  exuser,ok:=users[email]
	  if !ok{
		  fmt.Println("user not found")
		  w.Write([]byte("user was not found"))
		  return 
	  }

	  w.Header().Set("Content-type","application/json")
   json.NewEncoder(w).Encode(exuser)




}


func main(){
	 
	http.HandleFunc("/signup",senddetails)
	http.HandleFunc("/users",getdetails)
	http.HandleFunc("/users/",updatedetails)
    http.HandleFunc("/rm/",deletedetails)
	http.HandleFunc("/user/",getsingle)

	
    fmt.Println("server is starting at the port no 8080")
	http.ListenAndServe(":8080",nil)
	
}
