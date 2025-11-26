package main

import "fmt"

// func add(a int ,b int) int{
//     //  fmt.Println(a+b)
// 	 return a+b  
// }
// func add(a ,b int) int{
//     //  fmt.Println(a+b)
// 	 return a+b  
// }
// func add(a ,b int) (int,int){
//     //  fmt.Println(a+b)
// 	 return a+b  ,a-b
// }

func language()(string ,string, string){
	return "golang" ,"js" ,"java"
}




func main(){
// s,s1:=add(5,6)
// fmt.Println("the add of two " , s ,"the sub of two num",s1)

s,s1,_ :=language()
fmt.Println(s,s1)
}