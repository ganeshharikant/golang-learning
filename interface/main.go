package main

import "fmt"


type paymenter interface{
	pay(amount float32)
}
type payment struct{
	 gateway paymenter
}

func(p payment)makepayment(amount float32){
//   razorgw:=razorpay{}
//   razorgw.pay(amount)

    p.gateway.pay(amount)
}
type stripe struct{}
 
func( s stripe)pay(amount float32){
	fmt.Println("making payment using stripe",amount)
}


type razorpay struct{}

func(r razorpay)pay(amount float32){
	fmt.Println("making payment using rzaor pay",amount)
}

func main(){
	// stripegw:=stripe{}
newpay:=payment{
   gateway: stripe{},
}
newpay.makepayment(1000)
newpay1:=payment{
   gateway: razorpay{},
}
newpay1.makepayment(1000)



}


