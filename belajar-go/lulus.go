package main

import "fmt"

func main () {
	var x int
	fmt.Scan(&x)

	if x >= 60 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Gagal")
	}
}