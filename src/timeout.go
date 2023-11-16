package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(round int) int {
	last := 0
	for i := 0; i < round; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("sleep!")
	}
	return last + 1
}

func work(f func(int) int) chan int {
	c := make(chan int)
	var wg sync.WaitGroup

	ls := []int{3, 2, 5, 4, 3, 4, 4, 4, 4, 3, 3, 2, 2, 1}

	for _, i := range ls {
		wg.Add(1)

		go func(i int) {
			f(i)
			fmt.Println(i, "Done!")
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		c <- len(ls)
	}()

	return c
}

func runTimeOut() {
	c := work(myFunc)

	timer := time.NewTimer(4 * time.Second)
	defer timer.Stop()

	select {
	case done := <-c:
		fmt.Println("all done!", done)
	case <-timer.C:
		fmt.Println("time out!")
	}
}
