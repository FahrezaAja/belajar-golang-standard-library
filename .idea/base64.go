package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	value := "Fahreza Mandala Putra"

	encode := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded:", encode)

	decoded, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	} else {
		fmt.Println(string(decoded))
	}

}