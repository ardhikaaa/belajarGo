package main

import "fmt"

func main(){
	var mil float64

	fmt.Print("Masukan nilai mil:")
	fmt.Scan(&mil)

	km := mil * 1.6

	fmt.Printf("%g X 1.6 = %.1f", mil, km)
}