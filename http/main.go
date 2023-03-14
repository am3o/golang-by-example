package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Name  string  `json:"Foo"`
	Price float32 `json:"Bar"`
}

var products = []Product{
	{
		Name:  "bAnAnA",
		Price: 5,
	},
	{
		Name:  "ApPlE",
		Price: 1.0432,
	},
}

func Event(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(writer, request)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(fmt.Sprintf("%s %s", request.Method, request.URL.Path))
		next.ServeHTTP(writer, request)
	})
}

func getProducts(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		err := json.NewEncoder(writer).Encode(products)
		if err != nil {
			panic(err)
		}
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/api/v1/products/", getProducts)

	m := Event(Logger(r))

	if err := http.ListenAndServe(":8080", m); err != nil {
		panic(err)
	}
}
