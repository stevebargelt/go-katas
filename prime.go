package main

import (
	"fmt"
	"math"
)

func sieve(n int) []int {
	list := make(map[int]bool)

	for i := 2; i <= n; i++ {
		list[i] = true
	}

	limit := math.Sqrt(float64(n))

	for i := 2; i <= int(limit); i++ {
		for x := i + i; x <= n; x += i {
			list[x] = false
		}
	}

	output := make([]int, 1)
	for i := 2; i <= n; i++ {
		if list[i] {
			output = append(output, i)
		}
	}

	return output
}

func main() {

	fmt.Println("Hello from Prime")
	fmt.Println(sieve(2000))
}
