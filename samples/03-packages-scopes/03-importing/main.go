package main

import "fmt"
import f "fmt" // rename to a different variable.

func main() {
	fmt.Println("Hello!")
	f.Println("There!")
}