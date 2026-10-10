package main

import "fmt"

func main(){
	var y, x int
	
	fmt.Print("Masukan Jumlah Kue: ")
	fmt.Scan(&y)
	fmt.Print("Masukan Jumlah Keluarga: ")
	fmt.Scan(&x)

	sisa := y % x

	fmt.Printf("y = %d, x = %d -> %d %% %d = %d", y, x, y, x, sisa)
}