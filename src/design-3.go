package main

import "fmt"

func RunCheckDesign3() {
	total := 0
	fmt.Print("design-2...")
	startForDesign3(idCheck, make(chan int))
	println("total time cost", total)
}

func startForDesign3(workFunction func() int, next <-chan int) {
	workFunction()
	c := <-next
	fmt.Println(c)
}
