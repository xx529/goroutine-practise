package main

import "fmt"

func RunCheckDesign2() {
	fmt.Print("design-2...")

	passengerChannelList := make(chan struct{})

	var allChannelList []<-chan int

	for i := 0; i < NumOfChannel; i++ {
		c := startForDesign2(RunAllCheck, passengerChannelList)
		allChannelList = append(allChannelList, c)
	}

	for i := 0; i < Passengers; i++ {
		passengerChannelList <- struct{}{}
	}
	close(passengerChannelList)

	println("total time cost", cost(allChannelList))

}

func startForDesign2(workFunction func() int, queue <-chan struct{}) <-chan int {
	c := make(chan int)

	go func() {
		total := 0

		for {
			_, ok := <-queue
			if !ok {
				c <- total
				return
			}
			total += workFunction()
		}
	}()
	return c
}
