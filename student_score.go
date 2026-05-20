package main

import "fmt"

func main() {
	mahasiswa := map[string]int{
		"Ahmad":  90,
		"Syarif": 70,
		"Daniel": 45,
	}

	for nama, nilai := range mahasiswa {

		if nilai > 75 {
			fmt.Println(nama, "-", nilai, "-", "Lulus")
		} else {
			fmt.Println(nama, "-", nilai, "-", "Tidak lulus")
		}
	}
}
