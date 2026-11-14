package main

import "fmt"

func nilaiHuruf(x int) string {
	if x > 80 {
		return "A"
	} else if x > 70 {
		return "AB"
	} else if x > 65  {
		return "B"
	} else if x > 60 {
		return "BC"
	} else if x > 50 {
		return "C"
	} else if x > 40 {
		return "D"
	} else {
		return "E"
	}
}

func avg(total int, n int) float64 {

    if n == 0 {
        return 0.0
    }

    return float64(total) / float64(n)
}

func main() {
	var y, x, total int
	var a string
	var Avg float64
	fmt.Scan(&y)

	total = 0
	for i := 1; i <= y; i++ {
		fmt.Scan(&a, &x)
		nilai := nilaiHuruf(x)
		fmt.Printf("%s, %d, %s\n", a, x, nilai)
		total += x
		Avg =  avg(total, y)
	}

	fmt.Printf("%.2f", Avg)

}