package main

import "fmt"

func ifElse(i float64) {
	if i < 0 {
		fmt.Println("Negatif")
	} else if i == 0 {
		fmt.Println("Nol")
	} else {
		fmt.Println("Positif")
	}

}

func main() {
	var x, y float64

	fmt.Print("Masukan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukan bilangan y: ")
	fmt.Scan(&y)

	ifElse(x)
	ifElse(y)
}