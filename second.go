package main
import "fmt"

func main(){
	x := 10

	changeValue(&x)
	fmt.Println(x)
	
}

func changeValue(x *int){
	*x = 7
}

