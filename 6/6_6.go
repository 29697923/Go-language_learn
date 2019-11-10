package main

import "fmt"

func main(){
	var cheeses=make(map[string]int,2)
/*	var cheeses=make(map[value]int,2) Mapping is map[Type of mapping character] */
	cheeses["A"]=0
	cheeses["B"]=1
	fmt.Println(cheeses["A"])
	fmt.Println(len(cheeses))          //character length
	fmt.Println(cheeses["B"])
	fmt.Println(cheeses["c"])
}
