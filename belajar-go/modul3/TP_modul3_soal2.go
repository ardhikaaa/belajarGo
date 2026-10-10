package main

import "fmt"

func main() {
	var nilai bool

	fmt.Print("Masukan nilai (true/false):")
	fmt.Scan(&nilai)

	fmt.Printf("Nilai boolean %t dibaca lalu dicetak kembali", nilai)
}