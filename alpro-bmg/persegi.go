package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		var sisi int

		fmt.Scan(&sisi)

		luas := sisi * sisi
		keliling := 4 * sisi

		fmt.Println(luas, keliling)
	}
}