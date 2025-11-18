package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Fahreza", "Mandala", "Putra"})
	_ = writer.Write([]string{"Sekar", "Ayu"})
	_ = writer.Write([]string{"Rezaldi", "Seza"})

	writer.Flush()
}