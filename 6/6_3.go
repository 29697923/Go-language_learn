package main

import(
	"fmt"
)

func main(){
	var s =make([]string,2)
	s[0]="---------"
	s=append(s,"hello","Mikai")
	fmt.Println(s[0])
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[2])
	fmt.Println(s[3])
}
