package main

import "fmt"

func main() {
	var x, y, z int

	fmt.Scan(&x, &y, &z)

	if x >= y && x >= z {
		fmt.Print(x)
	} else if y >= z {
		fmt.Print(y)
	} else {
		fmt.Print(z)
	}
}