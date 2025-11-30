package main

import (
	"fmt"
	"slices"
)

func takeinput()[] int{
  var n int
 fmt.Println("enter the size of the array")
  fmt.Scan(&n)
 fmt.Println("enter the elements of the array \n")
 nums:=make([]int ,n)
 for i:=0;i<n;i++{
	fmt.Scan(&nums[i])
 }
 return nums



}
func compare(nums[]int,nums2[]int)bool {
   return slices.Equal(nums,nums2)
	

}

func modify(nums[] int, index int,num int)bool{
	  if index>=len(nums) || index<0{
		 return false
	  }
		nums[index]=num
        return true
}
func createmap()map[string]string {
	 m:=make(map[string]string)
	 m["name"]="ganesh"
	 m["lname"]="harikant" 
	 return m

}


func main(){
  
	nums:=takeinput();
	nums2:=takeinput();
	fmt.Println(nums);
	fmt.Println(nums2)
   b:=compare(nums,nums2)
   if(b==true){
	fmt.Println("both slices are equal")
   }else{
	fmt.Println("both slices are not equal")
    }
   c:=modify(nums,100,2)
   if(c==true){
	fmt.Println("yes it is suitable to modify the array")
   }else{
	fmt.Println("yes it is not suitable to modify the array")

   }
   fmt.Println(nums)
    m:=createmap()
	fmt.Println(m)


}


