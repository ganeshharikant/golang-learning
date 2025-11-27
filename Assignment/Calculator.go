package main

import "fmt"


func operator(a string ) bool{
switch a{
case "+":fmt.Println("you have choosed addition")
        return true;

               
case  "-":fmt.Println("you have choosed subtraction")
          return true
case  "*"  :fmt.Println("you have choosed multiplication")
            return true
case   "/"  :fmt.Println("you have choosed division")
            return true
default:     fmt.Println("you have choosed  the  option which cannot be performed in this calculator")
            return false

}



}

func cal(op string,a int ,b int)(int ,error){
   switch op{
            case "+": 
			        return a+b,nil
            case "-":
				    return a-b,nil
			case "*":
				    return a*b,nil
			case "/":
				    if b==0{
						return 0,fmt.Errorf("you cannot divide by zero")
					} else{
                        return a/b ,nil
					}

   }
   return 0, fmt.Errorf("u have choosen some differnt operation")
     
}

func main(){
 for{
	  fmt.Println("enter the operation that u want to perform:")
	  fmt.Println(" + ->  for addition") 
	  fmt.Println(" - ->  for subtraction") 
	  fmt.Println(" * ->  for multiplication") 
	  fmt.Println(" /  -> for divison ") 
	    fmt.Println("******************************")
        var op string
	   fmt.Scanln(&op)


	   val:=operator(op)
	   if !val{
           continue;
	   }
      
	var a int;
	var b int;
		fmt.Println("enter the value of a");
		fmt.Scanln(&a)
		fmt.Println("enter the value of b");
		fmt.Scanln(&b)
         
         res,err:=cal(op,a,b);

		 if err !=nil{
			 fmt.Println(err)
		 }else {
            switch op {
            case "+":
                fmt.Println("Value after addition is:", res)
            case "-":
                fmt.Println("Value after subtraction is:", res)
            case "*":
                fmt.Println("Value after multiplication is:", res)
            case "/":
                fmt.Println("Value after division is:", res)
            }
        }



  
 }
  




}












