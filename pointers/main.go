package main

import "fmt"

func change(num *int){
	*num=1;


}

func main(){
   num:=5;
   fmt.Println("the value before changing the value of number" ,num);
   
   
   change(&num)
   fmt.Println("the value after changing the value of number" ,num);
  



}                     