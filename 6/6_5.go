package main

import "fmt"

func main(){
	var cheeses=make([]string,2)
	cheeses[0]="Mariolles"
	cheeses[1]="Epoisses de bourgogne"
	var smllyCheeses=make([]string,2)
/*"	var smllyCheeses=make([]string)                this is false */
	copy(smllyCheeses,cheeses)
	fmt.Println(smllyCheeses)
}
