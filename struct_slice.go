package main

import "fmt"

type Mhs struct {
	Nama  string
	Nilai int
}

func nilai_mhs() {
	daftarMhs := []Mhs{
		{"Ahmad", 90},
		{"Syarif", 65},
		{"Daniel", 45},
	}
	for _, mhs := range daftarMhs {
		fmt.Println(mhs.Nama, "-", mhs.Nilai)
	}
}
