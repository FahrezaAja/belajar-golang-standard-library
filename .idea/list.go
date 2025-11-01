package main
import (
	"container/list"
	"fmt"
)


func main(){
	data := list.New()

	data.PushBack("Fahreza")
	data.PushBack("Sekar")
	data.PushBack("Diriku")

	for ambil := data.Front(); ambil != nil; ambil = ambil.Next(){
		fmt.Println(ambil.Value)
	}
}