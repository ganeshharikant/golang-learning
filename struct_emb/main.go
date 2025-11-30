package main

import "fmt"

 type customer struct{
	 name string
	 id   int
 }

 type order struct{
	  orderid int
	  customer
 }


func main(){

myorder:=order{ 
	   orderid: 1,
	   customer: customer{
		        name:"ghanesh",
				id:1,
	   },
	
}

fmt.Println(myorder)
fmt.Println(myorder.name)



}



