package main

import "time"

func Working(timeCost int) {
	time.Sleep(time.Millisecond * time.Duration(timeCost))
}

func idCheck() int {
	Working(IdCheckTimeCost)
	return IdCheckTimeCost
}

func bodyCheck() int {
	Working(BodyCheckTimeCost)
	return BodyCheckTimeCost
}

func xRayCheck() int {
	Working(XRayCheckCost)
	return XRayCheckCost
}

func RunAllCheck() int {
	total := 0
	total += idCheck()
	total += bodyCheck()
	total += xRayCheck()
	return total
}
