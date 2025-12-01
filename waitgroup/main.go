package main

import (
	"fmt"
	"sync"
	"time"
)

// func sayhello(wg *sync.WaitGroup){
//      defer wg.Done()
// 	fmt.Println("hello func has started")
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("hello func has ended")

// }

// func sayhi(id int){

//    fmt.Println(id)

// }

// func sayhi(id int) {
// 	time.Sleep(2 * time.Second) // Simulate slow task
// 	fmt.Println("Worker:", id)
// }


func sayhi(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second) // Simulate slow task
	fmt.Println("Worker:", id)
}

func main(){
// 	var wg sync.WaitGroup
//   fmt.Println("main started")
//     wg.Add(1)   
// go sayhello(&wg)
//   wg.Wait()

// time:=time.start();
// // time.Sleep(2* time.Second)
// //   fmt.Println("main has  ended")
//   for i:=0;i<=10;i++{
// 	    sayhi(i)
//   }

// time





	// fmt.Println("Main started (No Goroutine) at:", time.Now().Format("15:04:05"))

	// for i := 0; i <= 10; i++ {
	// 	sayhi(i)
	// }

	// fmt.Println("Main ended (No Goroutine) at:", time.Now().Format("15:04:05"))


fmt.Println("Main started (With Goroutine) at:", time.Now().Format("15:04:05"))

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go sayhi(i, &wg)
	}

	wg.Wait() // wait for all goroutines to finish

	fmt.Println("Main ended (With Goroutine) at:", time.Now().Format("15:04:05"))

	
}

