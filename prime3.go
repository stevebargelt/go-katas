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
		for j := i * i; j <= n; j += i {
			list[j] := false
		}
	}

	output := make([]int, 1)

	for k := 2; k <= n; k++ {
		if list[k] {
			output = append(output, k)
		}
	}

	return output

}

func main() {

	fmt.Println("Hello from Prime3")
	fmt.Println(sieve(20000))
}
