package main
import "fmt"
func main(){
	ifelseDemo()
}

	var age int
	fmt.Scanln(&age)
	if age > 18{
		fmt.Println("Adult")
	}else{
		fmt.Println("Minor")
	}
}func ifelseDemo(){