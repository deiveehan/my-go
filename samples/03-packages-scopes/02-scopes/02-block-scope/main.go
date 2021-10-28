package main

import "fmt"

func nope() {
	const ok = true
	var hello = "Hello"
	_ = hello
}

func main()  {
	fmt.Println(hello) // Hello not visible here.
}