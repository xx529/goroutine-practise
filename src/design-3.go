package main

import (
	"time"
)

func RunCheckDesign3() {
	p := make(chan struct{}, Passengers)
	startAllChannel(p)
	startAllChannel(p)
	startAllChannel(p)

	time.Sleep(5 * time.Second)
	for i := 1; i < Passengers; i++ {
		p <- struct{}{}
	}

	time.Sleep(5 * time.Second)
	close(p)
	time.Sleep(50 * time.Second)
}

func startChannel(workFunction func() int, next chan struct{}) (chan struct{}, chan struct{}, chan int) {
	total := 0
	queue := make(chan struct{}, 10)
	quit := make(chan struct{})
	result := make(chan int)

	go func() {
		for {
			select {
			case <-quit:
				result <- total
				return
			case v := <-queue:
				total += workFunction()
				if next != nil {
					next <- v
				}
			}
		}

	}()

	return queue, quit, result
}

func startAllChannel(p chan struct{}) {
	go func() {
		queue3, quit3, result3 := startChannel(xRayCheck, nil)
		queue2, quit2, result2 := startChannel(bodyCheck, queue3)
		queue1, quit1, result1 := startChannel(idCheck, queue2)

		for {
			select {
			case v, ok := <-p:
				if !ok {
					close(quit1)
					close(quit2)
					close(quit3)
					total := cost([]<-chan int{result1, result2, result3})
					println("total time cost", total)
					return
				}
				queue1 <- v
			}
		}
	}()
}
