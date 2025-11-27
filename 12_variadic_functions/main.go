package main

import "fmt"

func add(nums...int) int{
	total:=0
    for _ ,val:=range nums{
      total+=val
	}
	return total
}


func main(){



nums:=[]int{1,2,3}
sum:=add(nums...)
fmt.Println(sum)


}
