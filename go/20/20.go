package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func factorial(n int) int {
// 	sum := 1
// 	for i := 1; i < n+1; i++ {
// 		sum *= i
// 	}
// 	return sum
// }

func main() {
	fac := "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000"
	facArray := strings.Split(fac, "")
	sum := 0
	for i := range facArray {
		num, _ := strconv.Atoi(facArray[i])
		sum += num
	}
	fmt.Println(sum)
}
