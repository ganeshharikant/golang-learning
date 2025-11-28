package main

import (
	"fmt"
	"time"
)
type order struct{
	id string
	amount float32
	status string
	createat time.Time
}


func main(){
 var order1 order
//  var order2 order
 order1=order{
	id:"1",
	amount:500.00, 
	status:"delivered",
	createat:time.Now(),
 }

 order2:=order{
	id:"3",
	amount:5000.99,
	status:"packed",
	createat: time.Now(),

 }

 fmt.Println(order1.id)
 fmt.Println("details of the first order" ,order1)
  order1.status="request to return"
   fmt.Println("details of the second order" ,order2)
    fmt.Println(" after changing details of the first order" ,order1)
}