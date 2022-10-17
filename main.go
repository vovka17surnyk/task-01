package main

import (
	"io"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {

	ttp.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
