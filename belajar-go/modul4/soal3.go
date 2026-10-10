package main 

import (
	"fmt"
	"strings"
)

func isKabisat(tahun int) bool {
	return tahun % 4 == 0 && tahun % 100 != 0 || tahun % 400 == 0
}

func kapital(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func bulanValid(b string) bool {
	switch b {
	case "Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des":
		return true
	} 
	return false
}

func main(){
	var tahun int
	var bulan string
	
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&tahun)
	fmt.Print("Masukkan Bulan: ")
	fmt.Scan(&bulan)

	var hari string
	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		hari = "31"
	case "Apr", "Jun", "Sep", "Nov":
		hari = "30"
	case "Feb":
		if isKabisat(tahun) {
			hari = "29"
		} else {
			hari = "28"
		}
	default:
		saran := kapital(bulan)
		if bulanValid(saran) {
		fmt.Printf("%q tidak valid (huruf besar/kecil berbeda dari format yang diminta, seharusnya %q)", bulan, saran)
		} else {
			fmt.Printf("%q tidak valid", bulan)
		}
	}

	fmt.Println(hari)
}