package main

import "fmt"

func crivoEratostenes(n int64) (int64, bool) { // isprime de O(sqrt(n))
	if n <= 1 {
		return n, false
	}
	if n <= 3 {
		return n, true
	}
	if n%2 == 0 || n%3 == 0 {
		return n, false
	}

	// Só testa até √n
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+int64(2)) == 0 {
			return n, false
		}
	}
	return n, true
}

func main() {
	var sum int64 = 0

	for i := range 2000000 {
		num, status := crivoEratostenes(int64(i))
		if status {
			sum += num
		}
	}
	fmt.Println(sum)
}
