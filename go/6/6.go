package main

import (
	"fmt"
	"math/big"
)

func main() {
	sumSquares := new(big.Int)
	squareOfTheSum := new(big.Int)
	for i := 1; i <= 100; i++ {
		num := big.NewInt(int64(i))
		square := new(big.Int).Mul(num, num)
		sumSquares.Add(sumSquares, square)
	}
	sum := new(big.Int)
	for i := 1; i <= 100; i++ {
		sum.Add(sum, big.NewInt(int64(i)))
	}
	squareOfTheSum.Mul(sum, sum)
	result := new(big.Int).Sub(squareOfTheSum, sumSquares)
	fmt.Println(result)
}
