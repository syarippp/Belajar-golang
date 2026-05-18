package main

import "fmt"

func tegur(nama string) {
	fmt.Println("Kenapa kecil nilai kamu", nama, "?")
}
func motivasi(nama string) {
	fmt.Println("Nilai kamu cukup", nama, "tapi bisa ditingkatkan lagi ya")
}
func puji(nama string) {
	fmt.Println("Bagus sekali, pertahankan ya", nama)
}

func main() {

	var nama string
	var nilai int

	fmt.Print("Masukkan nama:")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan nilai:")
	fmt.Scanln(&nilai)

	if nilai > 90 {
		puji(nama)
	} else if nilai >= 75 {
		motivasi(nama)
	} else {
		tegur(nama)
	}
}
