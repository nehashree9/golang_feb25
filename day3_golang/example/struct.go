package main
import "fmt"
type student struct {
    Name string
    Number int
}
func main() {
    student1 := student{Name: "Alice", Number: 12345}
    st2 := student{Name: "Bob", Number: 67890}
    fmt.Println("Student 1:-","\nName:",student1.Name,"\nNumber:",student1.Number)
    fmt.Println("\nStudent 2:-", "\nName:",st2.Name,"\nNumber:",st2.Number)
}