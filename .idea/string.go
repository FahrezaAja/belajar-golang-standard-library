package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Fahreza Mandala Putra", "Fahreza"))
	fmt.Println(strings.Split("Fahreza Putra", " "))
	fmt.Println(strings.ToLower("Fahreza Mandala Putra"))
	fmt.Println(strings.ToUpper("Fahreza Mandala Putra"))
	fmt.Println(strings.Trim("        Fahreza        ", " "))
	fmt.Println(strings.ReplaceAll("Eja Eja Eja", "Eja", "Sekar"))
}