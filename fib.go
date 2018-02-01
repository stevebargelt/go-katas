package main

import (
	"fmt"
)

func fibFaster(x int) []int {

	sequence := make([]int, 2)
	sequence[0] = 0
	sequence[1] = 1
	for i := 2; i <= x; i++ {
		sequence = append(sequence, sequence[i-2]+sequence[i-1])
	}
	return sequence
}

var fibCount = 1

func calcFibAt(x int) int {

	fibCount++
	if x < 2 {
		return x
	}
	return calcFibAt(x-2) + calcFibAt(x-1)
}

func main() {

	for i := 0; i <= 40; i++ {
		fmt.Println("#", i, "Fib=", calcFibAt(i), "it took", fibCount, "calls to get here")

	}

	fmt.Println("\n\nand now for the fast Fib way:")
	fmt.Println(fibFaster(40))
}
