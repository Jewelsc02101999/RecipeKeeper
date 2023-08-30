package main

import (
	"codingforwomen/data"
	"codingforwomen/handlers"
	"fmt"
	"net/http"
)

func main() {
	data.FetchAllRecipes()
	fmt.Println(data.FetchAllRecipes)
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
}
