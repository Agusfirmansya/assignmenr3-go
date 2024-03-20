package main

import (
	"assignment_3/controllers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.CreateStatus)
	fmt.Println("localhost:8000")
	http.ListenAndServe(":8000", nil)
}