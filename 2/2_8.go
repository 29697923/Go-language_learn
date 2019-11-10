package main
import(
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	var b bool
	fmt.Println(b,reflect.TypeOf(b))
	var s_1 string
	b,s_1=strconv.ParseBool(false)
	fmt.Println(s_1,reflect.TypeOf(s_1))
}
