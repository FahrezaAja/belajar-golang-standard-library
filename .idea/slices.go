package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Fahreza", "Mandala", "Putra", "Sekar", "Ayu", "Rezaldi", "Seza"}
	number := []int{23, 20, 21}

	fmt.Println(slices.Min(number))
	fmt.Println(slices.Max(number))
	fmt.Println(slices.Contains(names, "Ayu"))
	fmt.Println(slices.Contains(names, "Budi"))
	fmt.Println(slices.Index(names, "Rezaldi"))
}