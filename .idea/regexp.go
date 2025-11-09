package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`e([a-z])a`)

	fmt.Println(regex.MatchString("eja"))
	fmt.Println(regex.MatchString("eza"))

	fmt.Println(regex.FindAllString("eja eza esa e2a era ela epa ", 10))

	
}