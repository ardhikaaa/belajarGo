package main

import "fmt"

func main() {
	var n int
	var hasil bool

	fmt.Print("nilai n: ")
	fmt.Scan(&n)

	hasil = n % 2 == 0

	fmt.Println(hasil)
}