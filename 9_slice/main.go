package main

import (
	"fmt"
	"slices"
)

func main(){
// var nums[]int


// nums[0]=1  // this is not possible  panic: runtime error: index out of range [0] with length 0

// nums=append(nums, 0)

// var nums=make([]int ,2,5)
// fmt.Println(cap(nums))
// nums=append(nums, 4)

// var nums=make([]int ,2,5)
// // 



// 

// var nums=[]int{1,2,3}

// nums=append(nums, 2);



// var nums=make([]int,0,5)
// var nums2=make([]int,len(nums))
// nums = append(nums, 2)
// copy(nums2,nums)

// var nums=[]int{1,2,3}
// fmt.Println(nums[0:2]) 
// fmt.Println(nums[:3]) 
// fmt.Println(nums[0:1]) 

// var nums2=[][]int{{2,4,4},{4,5,3}}
// fmt.Println(nums2) 


var nums1=[]int{1,2}
var nums2=[]int{1,2}
if slices.Equal(nums1,nums2){
	fmt.Println("both are equal")
}




// fmt.Println(nums2)



}


