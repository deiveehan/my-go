package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func responseSize(url string, channel chan int)  {
	fmt.Println("Getting", url)
	response, _:= http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	channel <- len(body)

}

func main() {
	sizes := make(chan int)
	// Runs in parallel, does not wait for one routing to complete processing.
	go responseSize("https://example.com/", sizes)
	go responseSize("https://golang.org/", sizes)
	go responseSize("https://golang.org/doc", sizes)
	fmt.Println(sizes)
	fmt.Println(sizes)
	fmt.Println(sizes)
}
