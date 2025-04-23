package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// log.Print("Request path:", r.URL.Path)
		log.Println("Hello World!")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oops!", http.StatusBadRequest)
			log.Printf("Error %s\n", err)
		} else {
			log.Printf("Data %s\n", d)
		}

		fmt.Fprintf(w, "Hello %s\n", d)

	})

	http.HandleFunc("/goodboy", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Stay connected!")
	})

	http.ListenAndServe(":9090", nil)
}
