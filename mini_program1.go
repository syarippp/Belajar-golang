package main

import "fmt"

type nilaiMhs struct {
	Nama   string
	Nilai  int
	Status string
}

var daftarMahasiswa []nilaiMhs

func hitungnilai(mhs nilaiMhs) string {

	if mhs.Nilai >= 75 {
		return "Lulus"
	}
	return "Tidak lulus"
}

func inputnilaiMhs() {

	var jumlah int

	fmt.Print("Jumlah Mahasiswa:")
	fmt.Scanln(&jumlah)

	for i := 0; i < jumlah; i++ {
		var mahasiswa nilaiMhs

		fmt.Print("Tambah Nama Mahasiswa:")
		fmt.Scanln(&mahasiswa.Nama)

		fmt.Print("Tambah Nilai Mahasiswa:")
		fmt.Scanln(&mahasiswa.Nilai)

		mahasiswa.Status = hitungnilai(mahasiswa)

		tampilkandataMhs(mahasiswa)

		fmt.Println()
	}
}

// func simpandataMhs(mahasiswa nilaiMhs) {

// }

func tampilkandataMhs(mahasiswa nilaiMhs) {
	daftarMahasiswa = append(daftarMahasiswa, mahasiswa)

	for _, mhs := range daftarMahasiswa {
		fmt.Println("==Data Nilai Mahasiswa==")
		fmt.Println("Nama:", mhs.Nama)
		fmt.Println("Nilai:", mhs.Nilai)
		fmt.Println("Status:", mhs.Status)

		fmt.Println()
	}
}
