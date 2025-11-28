package main

import "fmt"

// type order1 struct {
// 	id     string
// 	amount float32
// 	status string
// }



// func (o *order1)changestatus(a string ){
//        o.status=a

// }

// func (o order1)getamount()float32{
//   return o.amount 
// }

// func neworder(id string,amount float32,status string) order1{
// myorder:=order1{
//      id:id,
// 	 amount: amount,
// 	 status: status,
// }

// return myorder


// }

func main() {

	// order2 := order1{
	// 	id:     "1",
	// 	amount: 6000,
	// 	status: "packed",
	// }

	// fmt.Println(order2)



// order2.status="deliver"
// fmt.Println(order2)   


//  now i want to change the status of the order from the function of
// order2.changestatus("delivered")
// fmt.Println(order2)

// fmt.Println(order2.getamount())
// myorder:=neworder("1",1000,"deliever")
// myorder.changestatus("packed")
// fmt.Println(myorder)

lang:=struct{
	name string
	isgood bool
}{"golang",true}
fmt.Println(lang)




}