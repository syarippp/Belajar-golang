package main

import "fmt"

type nilaiMhs struct {
	Nama   string
	Nilai  int
	Status string
}

func inputnilaiMhs() []nilaiMhs {

	var jumlah int
	var daftarMahasiswa []nilaiMhs

	fmt.Print("Jumlah Mahasiswa:")
	fmt.Scanln(&jumlah)

	for i := 0; i < jumlah; i++ {

		var mahasiswa nilaiMhs

		fmt.Print("Tambah Nama Mahasiswa:")
		fmt.Scanln(&mahasiswa.Nama)

		fmt.Print("Tambah Nilai Mahasiswa:")
		fmt.Scanln(&mahasiswa.Nilai)

		mahasiswa.Status = hitungnilai(mahasiswa)

		daftarMahasiswa = append(daftarMahasiswa, mahasiswa)
	}
	return daftarMahasiswa
}

func hitungnilai(mhs nilaiMhs) string {

	if mhs.Nilai >= 75 {
		return "Lulus"
	}
	return "Tidak lulus"
}

func rataRata_nilai(daftarMahasiswa []nilaiMhs) float64 {
	var total int

	for _, mhs := range daftarMahasiswa {
		total += mhs.Nilai
	}

	rataRata := float64(total) / float64(len(daftarMahasiswa))

	return rataRata
}

func tampilkandataMhs(daftarMahasiswa []nilaiMhs) {

	fmt.Println("==Data Nilai Mahasiswa==")

	for _, mhs := range daftarMahasiswa {
		fmt.Println("Nama:", mhs.Nama)
		fmt.Println("Nilai:", mhs.Nilai)
		fmt.Println("Status:", mhs.Status)
		fmt.Println()
	}
	rata := rataRata_nilai(daftarMahasiswa)
	fmt.Println("Rata-Rata Semua Nilai:", rata)
}
