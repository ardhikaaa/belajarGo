package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	i := 1
	total := 0
	for i <= n {
		total += i
		i++ 
	}
	fmt.Println(total)
}