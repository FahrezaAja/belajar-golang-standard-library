package main

import (
	"fmt"
	"flag"
)

func main(){
	username := flag.String("username","root","database username")
	password := flag.String("password","root","database psword")
	host := flag.String("host","localhost","database host")
	port := flag.Int("port", 0 ,"database port")

	//jika ingin mengganti isi dari flag pada terminal, gunakan -username==nama untuk membuat perubahan setelah flag.go

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)

}