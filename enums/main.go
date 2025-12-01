package main

import "fmt"

type orderStatus string 

const(
   packed    orderStatus ="packed"
   delievered  ="delievered"
placed          ="placed"


)


func main(){

var  status orderStatus ="heelo"

// fmt.Println(orderStatus)
fmt.Println(status)

	
}
