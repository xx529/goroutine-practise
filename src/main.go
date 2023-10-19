package main

import (
	"fmt"
)

func main() {
	fmt.Println("total Passengers", Passengers)
	RunCheckDesign1()

	fmt.Println()
	RunCheckDesign2()

	fmt.Println()
	RunCheckDesign3()

	fmt.Println()
	runTimeOut()

	fmt.Println()
	runNotifyWait()

	fmt.Println()
	runClosePattern()

	fmt.Println()
	runPShutDown()
}
