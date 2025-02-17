package main

import (
	"crud-menggunakan-golang/config"
	"crud-menggunakan-golang/controller/categorycontroller"
	"crud-menggunakan-golang/controller/homcontroller"
	"crud-menggunakan-golang/controller/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//routes
	//homepage
	http.HandleFunc("/", homcontroller.Welcome)

	//categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/create", categorycontroller.Create)
	http.HandleFunc("/categories/update", categorycontroller.Update)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/create", productcontroller.Create)
	http.HandleFunc("/products/update", productcontroller.Update)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
