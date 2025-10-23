package main

import "fmt"

func main(){
	firstName := "Fahreza"
	lastName := "Putra"

	fmt.Println("Hello '", firstName, lastName,"'") //tapi ada spasi antara nama dan '
	//jika tidak ingin ada spasi, maka gunakan printf
	fmt.Printf("Hello '%s %s'\n",firstName, lastName) //maka tidak akan ada spasi antara nama dan '

	fmt.Printf("Hello '          %s       %s         ' \n",firstName, lastName) //bahkan kita dapat mengatur jaraknya
}