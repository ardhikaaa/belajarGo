package main

import "fmt"

func main() {
	var x int

	fmt.Print("Masukkan nilai Anda: ")
	fmt.Scan(&x)

	for x != -1 {
		if x >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}

		fmt.Print("Masukkan nilai Anda: ")
		fmt.Scan(&x)
	}

	fmt.Println("Program selesai.")
}