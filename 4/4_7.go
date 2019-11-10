package main
import "fmt"

func anuoTherFunction(f func())string{
	return f()
}
func main(){
	fn : =func() string{
		return "function called"
	}
	fmt.Println(anuoTherFunction(fn)
}
