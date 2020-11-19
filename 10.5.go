package main
import "fmt"

type A struct {
	fl float64
	int
	string
}

func main(){
	//var x A
	x:= &A{6.2,4,"assa"}
	fmt.Println(x)
}
