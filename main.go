package main

import (
	"io"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	7io.WriteString(w, "Hello world!7")
}

func main() {

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
