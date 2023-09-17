package main

import (
	"fmt"
)

func idCheck(pid int) int {
	Working(IdCheckTimeCost)
	fmt.Println(pid, "id check ok!")
	return IdCheckTimeCost
}

func bodyCheck(pid int) int {
	Working(BodyCheckTimeCost)
	fmt.Println(pid, "body check ok!")
	return BodyCheckTimeCost
}

func xRayCheck(pid int) int {
	Working(XRayCheckCost)
	fmt.Println(pid, "xray check ok!")
	return XRayCheckCost
}

func RunCheckDesign1(pid int) int {
	total := 0
	total += idCheck(pid)
	total += bodyCheck(pid)
	total += xRayCheck(pid)
	return total
}
