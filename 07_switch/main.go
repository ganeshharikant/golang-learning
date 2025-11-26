package main

import (
	"fmt"
)

func main() {
	// i := 1

	// switch i {
	// case 1:
	// 	fmt.Println("one");
	// case 2:
	// 	fmt.Println("two");
	// case 3:
	// 	fmt.Println("three");

	// }

	// switch i {
	// case "ganesh":
	// 	    fmt.Println("hello")
	// case "harikant":
	// 	    fmt.Println("harikant ganesh");		
		
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday,time.Sunday:
	// 	fmt.Println("its a weekednd")
	
		
	// }

whoami:=func (i interface{})  {
   switch t:=i.(type) {
   case int:
	    fmt.Println("integer");
   case string:
	    fmt.Println("string",t,"agnesh");
   }

	
}

whoami("ganesh");



}
