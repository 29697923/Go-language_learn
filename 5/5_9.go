package main

import "fmt"

func main(){
	number := []int{1,2,3,4}
	for i,n := range number{
		fmt.Println("The index of the loop is",i)
		fmt.Println("The value from the slice is",n)
	}
	str := []string{"a","b","v","d","e","f"}
	for s := range str{
		switch str{
		case "a":
			fmt.Printf("The indexx is %v",str)
		case "b":
			fmt.Printf("The indexx is %v",str)
		case "v":
			fmt.Printf("The indexx is %v",str)
		case "d":
			fmt.Printf("The indexx is %v",str)
		case "e":
			fmt.Printf("The indexx is %v",str)
		case "f":
			fmt.Printf("The indexx is %v",str)
		}
	}
}
