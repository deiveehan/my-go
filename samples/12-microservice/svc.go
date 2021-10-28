package main

import (
	"cmdspace.com/svc/handlers"
	"cmdspace.com/svc/models"
	"cmdspace.com/svc/repository"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	fmt.Println("Hello world")
	repository.AddProduct(models.Product{0, "Macbook pro", 2500})
	repository.AddProduct(models.Product{0, "Iphone", 1200})

	// Handler functions
	http.HandleFunc("/", handlers.HandleRequest)

	// Start a new server and on port 80
	err := http.ListenAndServe(":80", nil)

	// Error handling
	if err != nil {
		log.Println(err)
		return
	}

}