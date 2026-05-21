package main

import "fmt"

func jumlah_nilai(nilai1 int, nilai2 int) int {
	return nilai1 + nilai2
}

func lulus(nama string) {
	fmt.Println("Selamat anda lulus", nama)
}

func tidak_lulus(nama string) {
	fmt.Println("Anda tidak memenuhi minimal nilai", nama)
}

func main() {

	var nama string
	var nilai1 int
	var nilai2 int

	fmt.Print("Masukkan nama:")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan nilai 1:")
	fmt.Scanln(&nilai1)

	fmt.Print("Masukkan nilai 2:")
	fmt.Scanln(&nilai2)

	total_nilai := jumlah_nilai(nilai1, nilai2)

	if total_nilai >= 150 {
		lulus(nama)
	} else {
		tidak_lulus(nama)
	}
}
