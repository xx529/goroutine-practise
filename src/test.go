package main

import (
	"fmt"
	"time"
)

func work(round int) chan int {
	c := make(chan int)

	go func() {
		last := 0
		for i := 0; i < round; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("sleep!")
			last = i
		}
		c <- last + 1
	}()

	return c
}

func test() {
	c := work(5)
	t := <-c
	fmt.Println("round:", t)
}
