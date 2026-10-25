//KODE AWAL
// package main

// import "fmt"

// func main() {
// 	var a string
// 	var x, y, z, m, n int

// 	fmt.Scan(&a, &x)

// 	y = 1000
// 	z = 5000
// 	m = 1000
// 	n = 3000
// 	switch a {
// 	case "mobil":
// 		if x == 1 {
// 			fmt.Printf("Rp %d", z)
// 		} else {
// 			if x > 1 {
// 				fmt.Printf("Rp %d", z + n * x)
// 			}
// 		}
// 	case "motor":
// 		if x == 1 {
// 			fmt.Printf("Rp %d", y)
// 		} else {
// 			if x > 1 {
// 				fmt.Printf("Rp %d", y + m * x)
// 			}
// 		}
// 	default:
// 		fmt.Println("Kendaraan tidak valid")
// 	}
// }

//Perbaikan kode
package main

import "fmt"

func main() {
	var jenis string
	var jam, awal, lanjut int

	fmt.Scan(&jenis, &jam)

	switch jenis {
	case "motor":
		awal, lanjut = 2000, 1000
	case "mobil":
		awal, lanjut = 5000, 3000
	default:
		fmt.Println("Kendaraan tidak valid")
		return
	}

	if jam <= 0 {
		fmt.Println("Lama parkir tidak valid")
	} else {
		fmt.Printf("Rp%d\n", awal+(jam-1)*lanjut)
	}
}