package main

import (
	"fmt"
)

func Tian(i int){
	fmt.Println(&i)
	return
}

func main(){
	i :=1
	fmt.Println(&i)
	Tian(i)
}
