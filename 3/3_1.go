package main

import (
	"fmt"
	"reflect"
)

func main(){
	var s_1 string = "hello"
	var s_2 string = "mikai"
	fmt.Println(reflect.TypeOf(s_1))
	fmt.Println(s_1+s_2)
}
