package main

func main() {
	total := 0

	for i := 0; i < Passengers; i++ {
		total += RunCheckDesign1(i)
	}

	println("total time cost:", total)
}
