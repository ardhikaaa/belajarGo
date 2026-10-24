package main

import "fmt"

func main() {
	var x, y int
	var z string

	fmt.Scan(&x, &z, &y)

	switch z {
	case "+":
		fmt.Print(x + y)
	case "-":
		fmt.Print(x - y)
	case "*":
		fmt.Print(x * y)
	case "/":
		if y == 0 {
			fmt.Println("Tidak bisa dibagi nol")
		} else {
			fmt.Print(x / y)
		}
	default:
		fmt.Println("Operator tidak valid")
	}
}