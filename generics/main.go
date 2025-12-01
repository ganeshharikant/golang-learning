package main

import "fmt"


// func printanythingint(id[]int) {
// 	for _, v := range id {
// 		fmt.Println(v)
// 	}

///////////////////////////////////////////////////////////////////////////////////////////////////////

// }
// i wil be using generics now what wil i do is in order to print anything of integer and string i will be cretaing a singlke function to do it 

// type stack struct{
//    element [] int

// }
type stack[T any] struct{
   element [] T

}




// func printanything1[T comparable ](name [] T){
// 	for _,v :=range name{
// 		fmt.Println(v)
// 	}

// }





func main() {
	// names := []string{"ganesh", "supeerth"}
	// id:=[]int {1,2}
    // flag:=[] bool{true,false}

	// printanything(names)
	// printanythingint(id)
    
	// printanything1(names)
	// printanything1(id)
	// printanything1(flag)
    myele:=stack[int]{
        element:[]int{1,2,4},

	}
    myele1:=stack[string]{
        element:[]string{"ganesh" ,"harikant"},

	}

	fmt.Println(myele)
	fmt.Println(myele1)


}
