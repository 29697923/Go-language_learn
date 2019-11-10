package main

import (
	"fmt"
)

func tian(s *string){
	fmt.Println(s)
	return
}

func main(){
	s := "hello"
	fmt.Println(&s)
	tian(&s)

}
