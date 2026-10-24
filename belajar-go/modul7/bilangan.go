package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	if x == 0 {
		fmt.Println("Nol")
	} else {
		if x % 2 == 0 {
			fmt.Println("Genap")
		} else {
			fmt.Println("Ganjil")
		}
			if x > 0 {
			fmt.Println("positif")
		} else {
			fmt.Println("Negatif")
		}
	}
}