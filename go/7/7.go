package main

import "fmt"

// func isPrime(n int) bool { // isprime de O(sqrt(n))
// 	if n <= 1 {
// 		return false
// 	}
// 	if n <= 3 {
// 		return true
// 	}
// 	if n%2 == 0 || n%3 == 0 {
// 		return false
// 	}

// 	// Só testa até √n
// 	for i := 5; i*i <= n; i += 6 {
// 		if n%i == 0 || n%(i+2) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 1; i < n; i++ {
		if n%i == 0 && i != n && i != 1 {
			return false
		}
	}
	return true
}

func main() {
	primes := make([]int, 0, 120000)
	for i := range 200000 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes[10001])
}
