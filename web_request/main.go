package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning web api")
             res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
			 if(err!=nil){
				fmt.Println("Error getting res",err)
				return 
			 }

			 fmt.Printf("type of response: %T\n",res)
             data,err:=ioutil.ReadAll(res.Body)
			 if err!=nil{
				fmt.Println("error reading the resopnse",err)
				return 
			 }
			 fmt.Println("response",string(data))

}
