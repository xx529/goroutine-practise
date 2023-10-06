package main

import (
	"fmt"
	"time"
)

func spawn(f func(int)) chan string {
	quit := make(chan string)

	go func() {
		var job chan int
		for {
			select {
			case t := <-job:
				f(t)
			case <-quit:
				fmt.Println("do something before exit")
				quit <- "done"
			}
		}
	}()
	return quit
}

func runNotifyWait() {
	quit := spawn(busyWork)

	timer := time.NewTimer(4 * time.Second)
	defer timer.Stop()

	time.Sleep(2 * time.Second)
	quit <- "exit"

	select {
	case <-timer.C:
		fmt.Println("time out!")
	case status := <-quit:
		fmt.Println("status:", status)
	}
}
