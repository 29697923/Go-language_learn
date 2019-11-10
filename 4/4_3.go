package main

import "fmt"

func sumNumbers(number ...int)int{
	total:=0
	for _, number := rang number {
		total += number
	}
	return total
}


func main(){
	result := sumNumbers(1,2,3,4)
	fmt.Printf("The result is %v\n",result)
}
