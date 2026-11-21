package main
import (
	"fmt"
	"math"
	)
const pi float64 = 3.14

func persegiPanjang(p, l float64, luas, keliling *float64) {
	*luas = p * l
	*keliling = 2 * (p + l)
}

func lingkaran(r float64, luas, keliling *float64) {
	*luas = pi * r * r
	*keliling = 2 * pi * r 
}

func segitigaSiku(a, t float64, luas, keliling *float64) {
	*luas = a * t / 2
	*keliling = a + t + math.Sqrt(a * a + t * t)
}

func main(){
	var x string
	fmt.Scan(&x)

	switch x {
	case "P", "S":
		var a, b, luas, keliling float64
		fmt.Scan(&a, &b)
		if x == "P" {
			persegiPanjang(a, b, &luas, &keliling)
			fmt.Printf("Luas = %.2f\n", luas)
			fmt.Printf("Keliling = %.2f\n", keliling)
		} else {
			segitigaSiku(a, b, &luas, &keliling)
			fmt.Printf("Luas = %.2f\n", luas)
			fmt.Printf("Keliling = %.2f\n", keliling)
		}
	case "L":
		var r, luas, keliling float64
		fmt.Scan(&r)
		lingkaran(r, &luas, &keliling)
		fmt.Printf("Luas = %.2f\n", luas)
		fmt.Printf("Keliling = %.2f\n", keliling)
	default:
		fmt.Println("Kode bangun tidak valid")
	}
}