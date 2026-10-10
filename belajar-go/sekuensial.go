package main

import "fmt"

func main() {
	var namaDepan, namaBlkng, namaLengkap, inisial, usn, email string
	var tahun, panjangNama int
	var namaDpn bool

	fmt.Scan(&namaDepan)
	fmt.Scan(&namaBlkng)
	fmt.Scan(&tahun)

	namaLengkap = namaDepan + " " + namaBlkng
	// len = length; di baris ini menghitung jumlah karakter (byte) pada variabel namaLengkap
	panjangNama = len(namaLengkap)
	/* namaVariabel[low:high] = 
	low = karakter (byte) posisi awal yg ikut diambil
	high = karakter (byte) posisi akhir yg tidak ikut ambil
	contoh:
	nama := "ardhika"
	usn = nama[:2]
	output: "ar"
	*/
	inisial = namaDepan[:1] + namaBlkng[:1]
	/* namaBlkng[len(namaBlkng)-3:]
	mengambil 3 karakter (byte) belakang pada variabel
	contoh
	namaBlkng := "raditya"
	output: "tya"

	%02d
	2 = lebar minimal karakter
	0 = klo hasilnya krg dari 2 karakter maka akan ditambahkan angka 0 di depannya
	*/
	usn = namaDepan[:3] + namaBlkng[len(namaBlkng)-3:] + fmt.Sprintf("%02d", tahun % 100)
	email = namaDepan[:3] + "." + namaBlkng[len(namaBlkng)-3:] + "@student.telkomuniversity.ac.id"
	namaDpn = len(namaLengkap) >= 5

	fmt.Printf("%s\n", namaLengkap)
	fmt.Printf("%d\n", panjangNama)
	fmt.Printf("%s\n", inisial)
	fmt.Printf("%s\n", usn)
	fmt.Printf("%s\n", email)
	fmt.Printf("%t\n", namaDpn)
}