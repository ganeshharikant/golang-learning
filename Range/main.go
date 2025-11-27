package main

import "fmt"

func main(){

	// nums:=[]int{6,7,8}
	// for i:=0;i<len(nums);i++{
	// 	fmt.Println(nums[i])
	// }

// 	sum:=0;
// 	for _ ,v := range nums{
// 		sum=sum+v
// 	}
// 	fmt.Println(sum)
//   m:=map[string]string{"fname":"john","lname":"harikant"}
//   for _,v:= range m{
// 		fmt.Println(v)
	   
//   }


for i,c :=range"golang"{

fmt.Println(i,string(c))
// fmt.Println(i,(c)) //ascii value 
}
s2 := "a"
fmt.Println(len(s2))



}