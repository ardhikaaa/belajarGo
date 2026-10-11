package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	total := 0
	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
		total += i
	}
	fmt.Println()
	fmt.Println(total)

}