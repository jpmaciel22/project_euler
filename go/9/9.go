package main

import "fmt"

func main() {

	for i := 1; i < 1001; i++ {
		for j := 1; j < 1001; j++ {
			for k := 1; k < 1001; k++ {
				if k+i+j == 1000 && k*k+j*j == i*i {
					fmt.Println(k * i * j)
				}
			}
		}
	}
}
