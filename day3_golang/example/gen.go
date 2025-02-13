package main
import "fmt"
func main(){
    //ifelseDemo()
    //forCondiDemo()
     //forThreeVarDemo()
     //forPythonStyle()
mixed()
}
func forCondiDemo(){
    n := 1
    for n<5{
        n *= 2
    }
    fmt.Println(n)
}
func forThreeVarDemo(){
    sum:=0
    var i int
    for i=1;i<5;i++{
sum+=i
    }
    fmt.Println(sum)
}
func forPythonStyle(){
    strings:=[]string{"hello","world","go","2"}
    for i,s:=range strings{
        fmt.Println(i,s)
    }
}
func mixed(){
mixedSlice := []interface{}{"hello", 42, "world", 100}
    for _, v := range mixedSlice {
        fmt.Println(v)
    }
}