package  main

import "fmt"

func Mikai(portion int,eaten int)int{
	eaten=portion-eaten
	if eaten >= 5 {
		fmt.Printf("I'm full!I've eaten %d\n",eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry!I'%d\n",eaten)
	return Mikai(portion,eaten)
}

func main(){
	fmt.Println(Mikai(20,8))
}
