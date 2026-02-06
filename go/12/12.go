package main

import (
	"fmt"
	"math/big"
	"project_euler/libs"
)

func main() {
	sum := new(big.Int)
	current := big.NewInt(1)
	one := big.NewInt(1)

	for {
		sum.Add(sum, current)
		current.Add(current, one)
		if len(libs.Fatores(sum)) > 500 {
			fmt.Println(sum)
			break
		}
	}
}
