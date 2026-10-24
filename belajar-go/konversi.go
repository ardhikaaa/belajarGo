package main 

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	if x >= 85 && x <= 100 {
		fmt.Println("A")
	} else if x >= 70 && x <= 84 {
		fmt.Println("B")
	} else if x >= 55 && x <= 69 {
		fmt.Println("C")
	} else if x >= 40 && x <= 54 {
		fmt.Println("D")
	} else if x >= 0 && x <= 39 {
		fmt.Println("E")
	} else {
		fmt.Println("Nilai tidak valid")
	}
}