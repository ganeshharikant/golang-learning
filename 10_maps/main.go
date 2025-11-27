package main

import "fmt"

func main(){
 

	// m:=make(map[string]string)
	// m["name"]="golang"
	// fmt.Println(m["prices"])// here prices i not there in the map we will get the default value of the prices
    //   fmt.Print(len(m))


m1:= map[string]int {"price":30}
fmt.Println(m1)
m1["age"]=30
  fmt.Println(m1)
k,ok:=m1["price"]
 if ok{                
	fmt.Println("all ok")
 }else{
	fmt.Println("not done")
 }
 fmt.Println(k)


}

