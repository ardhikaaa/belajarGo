package main

import "fmt"

func main() {
	var x int
	var y string

	fmt.Scan(&x)
	fmt.Scan(&y)

	for i := 1; i <= x; i++ {
		fmt.Println(y)
	}
}