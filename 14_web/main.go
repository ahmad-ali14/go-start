package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Ahmad Ali</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1 style='color:red;' >I'm a Web Developer</h1>")

	})
	fmt.Println("server starting .....")
	http.ListenAndServe(":3000", nil)
}
