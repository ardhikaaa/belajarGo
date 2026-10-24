package main

import "fmt"

func main() {
	var x, y, z int 

	fmt.Scan(&x, &y, &z)

	if x <= 0 || x + y < z {
		fmt.Println("Bukan Segitiga")
	} else {
		if x == y && x == z && y == z {
			fmt.Println("Sama Sisi")
		} else if  x == y || x == z || y == z {
			fmt.Println("Sama Kaki")
		} else if x != y && x != z && y != z {
			fmt.Println("Sembarang")
		}
	}
}