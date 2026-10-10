package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("nilai a: ")
	fmt.Scan(&a)
	fmt.Print("nilai b: ")
	fmt.Scan(&b)

	fmt.Println(a > b, a == b, a < b)
}