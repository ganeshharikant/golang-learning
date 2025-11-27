package main

import "fmt"

func operator1(a string) bool {
	switch a {
	case "+":
		fmt.Println("you have choosed addition")
		return true

	case "-":
		fmt.Println("you have choosed subtraction")
		return true
	case "*":
		fmt.Println("you have choosed multiplication")
		return true
	case "/":
		fmt.Println("you have choosed division")
		return true
	default:
		fmt.Println("you have choosed  the  option which cannot be performed in this calculator")
		return false

	}

}

func cal1(op string, a int, b int) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("you cannot divide by zero")
		} else {
			return a / b, nil
		}

	}
	return 0, fmt.Errorf("u have choosen some differnt operation")

}

func main() {
	for {
		fmt.Println("enter the number")
		var a int
		fmt.Scanln(&a)
		for {
			fmt.Println("enter the operation that u want to perform:")
			fmt.Println(" + ->  for addition")
			fmt.Println(" - ->  for subtraction")
			fmt.Println(" * ->  for multiplication")
			fmt.Println(" /  -> for divison ")
			fmt.Println("******************************")
			var op string
			fmt.Scanln(&op)
			
			if op== "=" {
				fmt.Println("res", a)
				break;
			}
			val := operator1(op)

			if !val {
				continue
			}


			var b int

			fmt.Println("enter the number b")
			fmt.Scanln(&b)
      var err1 error;
			a , err1 = cal1(op, a, b)

			if err1 != nil {
				fmt.Println(err1)
				continue
			}
			fmt.Println("Current result:", a)

		}

	}

}
