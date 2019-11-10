package main

import "fmt"

func Mikai(b bool) bool{
	if b{
		b := false
		return b
	}else{
		b := true
		return b
	}
}

func main(){
	b:=Mikai(false)
	if b {
		fmt.Println("b is false")
	}else if b {
		fmt.Println("b is true")
	}
}
