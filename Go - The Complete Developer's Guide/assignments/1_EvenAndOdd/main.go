package main

import "fmt"

func main() {
	ints := []int{}
	for i := 0; i <= 10; i++ {
		ints = append(ints, i)
	}

	for _, val := range ints {
		if val%2 == 0 {
			fmt.Printf("%d is even.\n", val)
		} else {
			fmt.Printf("%d is odd.\n", val)
		}
	}
}
