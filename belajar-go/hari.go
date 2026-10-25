package main 

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	//KODE AWAL
	// switch x {
	// case 1:
	// 	fmt.Println("Januari (31 hari)")
	// case 2:
	// 	fmt.Println("Febuari (28 hari)")
	// case 3:
	// 	fmt.Println("Maret (31 hari)")
	// case 4:
	// 	fmt.Println("April (30 hari)")
	// case 5:
	// 	fmt.Println("Mei (31 hari)")
	// case 6:
	// 	fmt.Println("Juni (30 hari)")
	// case 7:
	// 	fmt.Println("Juli (31 hari)")
	// case 8:
	// 	fmt.Println("Agustus (31 hari)")
	// case 9:
	// 	fmt.Println("Spetember (30 hari)")
	// case 10:
	// 	fmt.Println("Oktober (31 hari)")
	// case 11:
	// 	fmt.Println("November (30 hari)")
	// case 12:
	// 	fmt.Println("Desember (31 hari)")
	// default:
	// 	fmt.Println("Bulan tidak valid")
	// }

	//Perbaikan kode
	var nama string
	var hari int

	switch x {
	case 1:
		nama, hari = "Januari", 31
	case 2:
		nama, hari = "Februari", 28
	case 3:
		nama, hari = "Maret", 31
	case 4:
		nama, hari = "April", 30
	case 5:
		nama, hari = "Mei", 31
	case 6:
		nama, hari = "Juni", 30
	case 7:
		nama, hari = "Juli", 31
	case 8:
		nama, hari = "Agustus", 31
	case 9:
		nama, hari = "September", 30
	case 10:
		nama, hari = "Oktober", 31
	case 11:
		nama, hari = "November", 30
	case 12:
		nama, hari = "Desember", 31
	default:
		fmt.Println("Bulan tidak valid")
		return
	}
	fmt.Printf("%s (%d hari)\n", nama, hari)
}

// pengelompokan
// switch x {
// case 4, 6, 9, 11:
// 	hari = 30
// case 2:
// 	hari = 28
// default:
// 	hari = 31
// }