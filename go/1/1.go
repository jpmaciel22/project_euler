package main

import "fmt"

func main() {
	num := 0
	for i := range 1000 {
		if i%3 == 0 || i%5 == 0 {
			num += i
		}
	}
	fmt.Println(num)
}
