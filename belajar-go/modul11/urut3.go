package main

import "fmt"

func tukar(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}

func urut3(a, b, c *int) {
	tukar(a, b)
	tukar(b, c)
	tukar(a, b)
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Printf("Sebelum: %d %d %d\n", a, b, c)
	urut3(&a, &b, &c)
	fmt.Printf("Sesudah: %d %d %d", a, b, c)
}