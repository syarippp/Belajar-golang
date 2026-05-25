package main

import "fmt"

type dataMhs struct {
	Nama   string
	Nilai  int
	Status string
}

func cekstatus(mhs dataMhs) string {

	if mhs.Nilai >= 75 {
		return "Lulus"
	}
	return "Tidak lulus"
}

func inputdataMhs() {

	var mahasiswa dataMhs

	fmt.Print("Masukkan data nama mahasiswa:")
	fmt.Scanln(&mahasiswa.Nama)

	fmt.Print("Masukkan data nilai mahasiswa:")
	fmt.Scanln(&mahasiswa.Nilai)

	mahasiswa.Status = cekstatus(mahasiswa)

	fmt.Println("Nama:", mahasiswa.Nama)
	fmt.Println("Nilai:", mahasiswa.Nilai)
	fmt.Print("Status:", mahasiswa.Status)
}
