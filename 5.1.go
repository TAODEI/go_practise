package main
import(
		"fmt"
		"strconv"
	  )
func main(){
	orig :="ABC"
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an,err := strconv.Atoi(orig)
	if err != nil{
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an += 5
	news := strconv.Itoa(an)
	fmt.Printf("The new string is:%s\n",news)
}


