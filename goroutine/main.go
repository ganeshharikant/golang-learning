package main

import (
	"fmt"
	"time"
)

func task(i int) {
	fmt.Println(i);

}

func main() {
	for i := 0; i <= 10; i++ {
		// go task(i)
      go func(i int ){
        fmt.Println(i)
	  }(i)

	}
  time.Sleep(2*time.Second)
}