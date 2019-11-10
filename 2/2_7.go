package main
import(
	"fmt"
	"reflect"
)
func main(){
	var a int=19
	var b float64=.5
	var c bool=true
	var d string="miku"


	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
}
