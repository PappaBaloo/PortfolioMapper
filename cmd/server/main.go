package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Server is running...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Nigger!")
	})
	http.ListenAndServe(":8080", nil)
}
