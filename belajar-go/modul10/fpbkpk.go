package main

import "fmt"

func fpb(x, y int) int {
	for y != 0 {
		temp := y
		y = x%y
		x = temp
	}
	return x
}

func kpk(x, y int) int {
	return x / fpb(x, y) * y
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fpb := fpb(x, y)
	kpk := kpk(x, y)

	fmt.Println(fpb)
	fmt.Println(kpk)
}