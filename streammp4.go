package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./Video/"))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Println("Live on 8080: ")
	http.ListenAndServe(":8080", nil)
}
