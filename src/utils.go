package main

import "time"

func Working(timeCost int) {
	time.Sleep(time.Millisecond * time.Duration(timeCost))
}
