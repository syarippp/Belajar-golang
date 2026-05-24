package main

import "fmt"

func sapa(nama string) {
	fmt.Println("Halo", nama)
}
func panggil() {
	sapa("Ahmad")
	sapa("Syarif")
}
