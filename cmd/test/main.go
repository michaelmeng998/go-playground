package main

import (
	"fmt"
	"io"
	"os"
)

//...
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Greetings, %s", name)
}

//...
func Sum(x int, y int) int {
	return x + y
}

func main() {
	Greet(os.Stdout, "Michael \n")

	sum := Sum(5, 5)

	//print sum to console

	fmt.Printf("sum is %d \n", sum)
}
