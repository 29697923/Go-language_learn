package main

import "fmt"

func main(){
	i:=3
	switch i {
	case 1:
		fmt.Println("one")
		i=i+1
	case 2:
		fmt.Println("two")
		i=i-2
	case 3:
		fmt.Println("three")
		i=i-1
	}

}
