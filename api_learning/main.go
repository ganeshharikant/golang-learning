package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r* http.Request){
       fmt.Fprintln(w,"heelo   ganesh")

}

func main() {
 n 
	http.HandleFunc("/",handler)
	fmt.Println("i am listening to this port")
	http.ListenAndServe(":8080", nil) 
}

