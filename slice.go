package main

import (
	"fmt"
	"strings"
)

func main() {
	nama := []string{"Ahmad", "Syarif", "Daniel"}

	fmt.Println(strings.Join(nama, ","))
}
