package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci1(n int, check *int) {
	var next int
	var first , second int = 0,1
	for i := 0; i < n; i++ {
		if (i <= 1){
			next = i
		}else {
			next = first + second
			first = second
			second = next
		}
		fmt.Printf("%d   ", next);
	}
	fmt.Println()
	*check = 1

}



func Fibonacci2(n int, wg *sync.WaitGroup) {
	var next int
	var first , second int = 0,1
	for i := 0; i < n; i++ {
		if (i <= 1){
			next = i
		}else {
			next = first + second
			first = second
			second = next
		}
		fmt.Printf("%d   ", next);
	}
	fmt.Println()
	wg.Done()

}

func Fibonacci3(n int) {
	var next int
	var first , second int = 0,1
	for i := 0; i < n; i++ {
		if (i <= 1){
			next = i
		}else {
			next = first + second
			first = second
			second = next
		}
		fmt.Printf("%d   ", next);
	}
	fmt.Println()


}

func concurrencyMain() {
	fmt.Printf("First 10 Fibonacci series are (using local varaable and sleep):\n");
	var mycheck int = 0
	go Fibonacci1(10,&mycheck)
	for mycheck==0{
		time.Sleep(time.Microsecond)
	}


	fmt.Printf("First 20 Fibonacci series are (use of WaitGroup):\n");
	wg := sync.WaitGroup{}
	wg.Add(1)
	go Fibonacci2(20,&wg)
	wg.Wait()

	fmt.Printf("First 30 Fibonacci series are:\n");
	go Fibonacci3(30)
	time.Sleep(5*time.Microsecond)

	//runtime.Breakpoint()

	fmt.Printf("End of Fibonacci series \n");



	//fmt.Println("GoMaxProcess",runtime.GOMAXPROCS(1))
}