package main
import "fmt"
func main()  {
	x := make([]int, 5, 10)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 100, 200,40,30,32)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}