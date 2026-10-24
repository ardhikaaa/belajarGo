package main

import "fmt"

func main () {
	var x int
	fmt.Scan(&x)

	if x % 4 == 0 && x % 100 != 0 || x % 400 == 0 {
		fmt.Println("Kabisat")
	} else {
		fmt.Println("Bukan Kabisat")
	}
}