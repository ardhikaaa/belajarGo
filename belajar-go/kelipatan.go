package main

import "fmt"

func main() {
	var n, a, b int
	var hasil bool

	fmt.Print("nilai n, a, b: ")
	fmt.Scan(&n, &a, &b)
	/*fmt.Print("nilai a: ")
	fmt.Scan(&a)
	fmt.Print("nilai b: ")
	fmt.Scan(&b)*/

	hasil = (n % a == 0) && (n % b == 0)

	fmt.Println(hasil)
}