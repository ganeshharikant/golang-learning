package main

import "fmt"

func main() {
	i:=1
	for i<=3{
		if i==2{
			i++;
			continue;
		}
		fmt.Println(i);
		i++;
	}
//    for {

//    }   

// for i:=0;i<=3;i++{
// 	fmt.Println(i);
// }

for i:=range 3{
	fmt.Println(i);

}



}
