package main
import "fmt"
func main(){
	for i:=0;i<=100;i++{
		switch{
			case i%3==0&&i%5!=0:
				fmt.Println("Fizz")
			case i%3!=0&&i%5==0:
	            fmt.Println("Buzz")
			case i%3==0&&i%5==0:
                fmt.Println("FizzBuzz")
			default:
				fmt.Println(i)
		}
	}
}
