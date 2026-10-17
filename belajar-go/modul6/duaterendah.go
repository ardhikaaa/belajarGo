package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var nama1, nama2 string 
	var nilai1, nilai2 int

	i := 0
	for i < n { 
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if i == 0 {
			nama1, nilai1 = nama, nilai
		} else if i == 1 {
			if nilai < nilai1 {
				nama2, nilai2 = nama1, nilai1
				nama1, nilai1 = nama, nilai
			} else {
				nama2, nilai2 = nama, nilai
			}
		} else {
			if nilai < nilai1 {
				nama2, nilai2 = nama1, nilai1
				nama1, nilai1 = nama, nilai
			} else if nilai < nilai2 {
				nama2, nilai2 = nama, nilai
			}
		}
		i++
	}

	fmt.Println(nama1)
	fmt.Println(nama2)
}