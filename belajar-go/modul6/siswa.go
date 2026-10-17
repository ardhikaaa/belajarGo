package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	for n >= 60 {
		fmt.Println("Anda Lulus Ujian")
		break
	}

	for n < 60 && n >= 0 {
		fmt.Print("Anda Tidak Lulus Ujian")
		break
	}

	for n < 0 {
		fmt.Println("Program Selesai")
		break
	}
}