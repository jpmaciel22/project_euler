package main

import "fmt"

func main() {
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	sum := 0
	i := 2
	for {
		lasti := i - 1
		lastlasti := i - 2
		cache[i] = cache[lasti] + cache[lastlasti]
		if cache[i] > 4000000 {
			break
		}
		i++
	}
	for i := range cache {
		if cache[i]%2 == 0 {
			sum += cache[i]
		}
	}
	fmt.Println(sum)
}
