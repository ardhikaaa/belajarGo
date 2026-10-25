package main 

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Println("Januari (31 hari)")
	case 2:
		fmt.Println("Febuari (28 hari)")
	case 3:
		fmt.Println("Maret (31 hari)")
	case 4:
		fmt.Println("April (30 hari)")
	case 5:
		fmt.Println("Mei (31 hari)")
	case 6:
		fmt.Println("Juni (30 hari)")
	case 7:
		fmt.Println("Juli (31 hari)")
	case 8:
		fmt.Println("Agustus (31 hari)")
	case 9:
		fmt.Println("Spetember (30 hari)")
	case 10:
		fmt.Println("Oktober (31 hari)")
	case 11:
		fmt.Println("November (30 hari)")
	case 12:
		fmt.Println("Desember (31 hari)")
	default:
		fmt.Println("Bulan tidak valid")
	}
}