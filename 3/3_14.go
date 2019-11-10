package main

import(
	"fmt"
)

const s string="Hello,world"
const i int   =11

func tian(i *int){
	fmt.Println(i)
}

func main(){
	fmt.Println(s)
	tian(&i)
}
