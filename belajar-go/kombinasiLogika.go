package main

import "fmt"

func main() {
	var p, q int
	var hasil1, hasil2, hasil3 bool

	fmt.Print("nilai p, q: ")
	fmt.Scan(&p, &q)

	hasil1, hasil2, hasil3 =  (p % 2 == 0) || (q % 2 == 0), (p % 2 != 0) && (q % 2 != 0), !(p == q)

	fmt.Println(hasil1, hasil2, hasil3)
}