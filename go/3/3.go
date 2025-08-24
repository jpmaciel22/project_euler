package main

import "fmt"

func main() {
	n := 600851475143
	largestFactor := 1

	for n%2 == 0 {
		largestFactor = 2
		n = n / 2
		fmt.Printf("Fator encontrado: 2, número restante: %d\n", n)
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			largestFactor = i
			n = n / i
			fmt.Printf("Fator encontrado: %d, número restante: %d\n", i, n)
		}
	}

	if n > 1 {
		largestFactor = n
		fmt.Printf("Último fator primo: %d\n", n)
	}

	fmt.Printf("\nResposta: %d\n", largestFactor)
}
