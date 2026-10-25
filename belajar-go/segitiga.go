package main

import "fmt"

func main() {
	var x, y, z int 

	fmt.Scan(&x, &y, &z)

	//KODE AWAL
	// if x <= 0 || x + y < z {
	// 	fmt.Println("Bukan Segitiga")
	// } else {
	// 	if x == y && x == z && y == z {
	// 		fmt.Println("Sama Sisi")
	// 	} else if  x == y || x == z || y == z {
	// 		fmt.Println("Sama Kaki")
	// 	} else if x != y && x != z && y != z {
	// 		fmt.Println("Sembarang")
	// 	}
	// }

	//Perbaikan kode 
	if x <= 0 || y <= 0 || z <= 0 || x+y <= z || x+z <= y || y+z <= x {
		fmt.Println("Bukan segitiga")
	} else if x == y && y == z {
		fmt.Println("Sama sisi")
	} else if x == y || x == z || y == z {
		fmt.Println("Sama kaki")
	} else {
		fmt.Println("Sembarang")
	}
}