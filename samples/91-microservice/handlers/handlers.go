package handlers

import (
	"cmdspace.com/svc/models"
	"cmdspace.com/svc/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received: ", r.Method)

	switch r.Method {
	case http.MethodGet:
		list(w, r)
	case http.MethodPost:
		add(w, r)
	default:
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		break
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)
	var product models.Product
	err := json.Unmarshal(payload, &product)

	if err != nil || product.Name == "" || product.UnitPrice ==0 {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}
	product.ID = repository.AddProduct(product)

	w.WriteHeader(201)
	json, _ := json.Marshal(product)
	w.Write(json)
	log.Println("Response returned: ", 201)
}

func list(w http.ResponseWriter, r *http.Request) {
	products := repository.GetProducts()
	json, _ := json.Marshal(products)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)

	log.Println("Response returned:", 200)
}