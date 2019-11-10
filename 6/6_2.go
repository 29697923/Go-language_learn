package main

import(
	"fmt"
)

func main(){
	var cheeses =make([]string,2)
	cheeses[0]="Hello"
	cheeses[1]="Word"
	cheeses=append(cheeses,"Mikai")
	fmt.Println(cheeses)
}
