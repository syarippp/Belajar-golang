package main

import "fmt"

type Mahasiswa struct {
	Nama     string
	Umur     int
	Semester int
}

func main() {
	mhs1 := Mahasiswa{
		Nama:     "Ahmad",
		Umur:     22,
		Semester: 8,
	}
	fmt.Println("Nama:", mhs1.Nama)
	fmt.Println("Umur:", mhs1.Umur)
	fmt.Println("Semester:", mhs1.Semester)
}
