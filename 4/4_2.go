package main

import (
	"fmt"
)

func Mikai_2()(int,string){
	i_1:=20
	s_1:="string_1"
	return i_1,s_1
}

func Mikai_1(i_2 int,s_2 string)(int,string){
	return i_2,s_2
}

func main(){
	var i,s=Mikai_1(10,"string_2")
	fmt.Printf("%v %v\n",i,s)
	i,s=Mikai_2()
	fmt.Printf("%v %v\n",i,s)
}
