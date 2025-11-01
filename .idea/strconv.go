package main

import (
	"strconv"
	"fmt"
)

func main(){

	boolean, err := strconv.ParseBool("true")
	if err == nil{
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	resulsint, err := strconv.Atoi("1000")
	if err == nil{
		fmt.Println(resulsint)
	} else {
		fmt.Println("Error", err.Error())
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	stringInt := strconv.Itoa(00000)
	fmt.Println(stringInt)
}