package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu (Akhir Pekan)")
	case 7:
		fmt.Println("Minggu (Akhir Pekan)")
	default: 
	fmt.Println("Nomor hari tidak valid")
	}
}