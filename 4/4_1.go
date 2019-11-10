package main
import(
	"fmt"
)

func addUp(x int,y int)int{
	return x+y;
}

func main(){
	i := addUp(10,50)
	fmt.Println(i)
}
