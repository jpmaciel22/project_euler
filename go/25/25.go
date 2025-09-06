package main

import (
	"fmt"
	"project_euler/libs"
)

func main() {

	first1000digit, index := libs.FibonacciNoRecursion(1000)
	fmt.Println(first1000digit, index)

}
