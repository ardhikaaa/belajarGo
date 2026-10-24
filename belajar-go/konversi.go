package main 

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	switch {
	case x >= 85 && x <= 100:
		fmt.Println("A")
	case x >= 70 && x <= 84:
		fmt.Println("B")
	case x >= 55 && x <= 69:
		fmt.Println("C")
	case x >= 40 && x <= 54:
		fmt.Println("D")
	case x >= 0 && x <= 39:
		fmt.Println("E")
	default:
		fmt.Println("Operator tidak valid")
	}
}