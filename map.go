package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"nama": "ahmad",
		"asal": "malang",
	}

	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa["nama"])
	fmt.Println(mahasiswa["asal"])
}
