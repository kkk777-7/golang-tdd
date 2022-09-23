package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(b io.Writer, word string) {
	fmt.Fprintf(b, "Hello, %s", word)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
