package main

import (
	"fmt"
)

func main(){
	var play=make(map[string]int)
	play["A"]=1
	fmt.Println(play)
	play["B"]=2
	fmt.Println(play)
	play["C"]=3
	fmt.Println(play)
	delete(play,"C")
	fmt.Println(play)
}
