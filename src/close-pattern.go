package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething() chan struct{} {
	quit := make(chan struct{})
	job := make(chan int, 10)
	wg := sync.WaitGroup{}

	for k := 2; k < 10; k++ {
		job <- k
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("worker", i, "started")
			for {
				j, ok := <-job
				if !ok {
					fmt.Println("job channel closed")
					return
				}
				fmt.Println("worker", i, "received", j)
				busyWork(j)
			}
		}(i)
	}

	go func() {
		<-quit
		close(job)
		wg.Wait()
		quit <- struct{}{}
	}()

	return quit
}

func runClosePattern() {
	quit := doSomething()

	time.Sleep(2 * time.Second)
	quit <- struct{}{}

	timer := time.NewTimer(3 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("time out!")
	case status := <-quit:
		fmt.Println("status:", status)
	}
}
