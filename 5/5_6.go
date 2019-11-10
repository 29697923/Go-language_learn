package main

import "fmt"

func main(){
	s:="s"
	switch s{
	case "a":
		fmt.Println("The value is 'a'")
	case "b":
		fmt.Println("The value is 's'")
	default:
		fmt.Println("The value is default")
	}
}
