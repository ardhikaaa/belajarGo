package main

import "fmt"

func main() {
	var x, y int
	var diskon, member float64

	fmt.Scan(&x, &y)

	switch y {
	case 1:
		member = 0.05
		if x >= 500000 {
			diskon = 0.15
			fmt.Print(float64(x) - (float64(x) * diskon) - (float64(x) * member))
		} else if x >= 200000 {
			diskon = 0.10
			fmt.Print(float64(x) - (float64(x) * diskon) - (float64(x) * member))
		} else if x >= 100000 {
			diskon = 0.05
			fmt.Print(float64(x) - (float64(x) * diskon) - (float64(x) * member))
		} else {
			fmt.Print(float64(x) - (float64(x) * member))
		}
	case 0:
		if x >= 500000 {
			diskon = 0.15
			fmt.Print(float64(x) - (float64(x) * diskon))
		} else if x >= 200000 {
			diskon = 0.10
			fmt.Print(float64(x) - (float64(x) * diskon))
		} else if x >= 100000 {
			diskon = 0.05
			fmt.Print(float64(x) - (float64(x) * diskon))
		} else {
			fmt.Print(x)
		}
	}
}