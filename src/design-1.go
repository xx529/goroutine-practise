package main

import "fmt"

func RunCheckDesign1() {
	total := 0
	fmt.Print("design-1...")
	for i := 0; i < Passengers; i++ {
		total += RunAllCheck()
	}
	println("total time cost", total)
}
