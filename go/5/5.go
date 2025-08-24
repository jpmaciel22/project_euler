package main

import "fmt"

func main() {
	smallestDivisbleNoRemainder := 0
	for i := 20; i <= 1000000000000000000; i++ {
		divisbleByAll := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				divisbleByAll = false
				break
			}
		}
		if divisbleByAll {
			smallestDivisbleNoRemainder = i
			break
		}
	}

	fmt.Println(smallestDivisbleNoRemainder)
}
