package main

import (
	"fmt"
	"strings"
)

func slice_data() {
	nama := []string{"Ahmad", "Syarif", "Daniel"}

	fmt.Println(strings.Join(nama, ","))
}
